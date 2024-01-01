package gopilot

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
	"github.com/webteleport/utils"
	"github.com/webteleport/webteleport/ufo"
)

const tokenUrl = "https://api.github.com/copilot_internal/v2/token"
const completionsUrl = "https://api.githubcopilot.com/chat/completions"

type Model struct {
	ID      string  `json:"id"`
	Object  string  `json:"object"`
	Created int     `json:"created"`
	OwnedBy string  `json:"owned_by"`
	Root    string  `json:"root"`
	Parent  *string `json:"parent"`
}

type ModelList struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

var ghuToken = ""

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func Run(args []string) error {
	handler := copilotHandler()
	handler = utils.WellKnownHealthMiddleware(utils.GzipMiddleware(handler))
	arg0 := Arg0(args, "https://ufo.k0s.io")
	if arg0 == "local" {
		port := utils.EnvPort(":8000")
		log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
		return http.ListenAndServe(port, handler)
	}
	return ufo.Serve(arg0, handler)
}

func copilotHandler() http.Handler {
	err := godotenv.Load()
	if err == nil {
		ghuToken = os.Getenv("GHU_TOKEN")
	}

	log.Printf("Server is running with GHU_TOKEN=%s", ghuToken)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
		curl --location 'http://127.0.0.1:8081/v1/chat/completions' \
		--header 'Content-Type: application/json' \
		--header 'Authorization: Bearer ghu_xxx' \
		--data '{
		  "model": "gpt-4",
		  "messages": [{"role": "user", "content": "hi"}]
		}'`)
	})

	r.GET("/v1/models", func(c *gin.Context) {
		c.JSON(http.StatusOK, models())
	})

	r.POST("/v1/chat/completions", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, must-revalidate")
		c.Header("Connection", "keep-alive")

		forwardRequest(c)
	})

	return r
}

func forwardRequest(c *gin.Context) {
	var jsonBody map[string]interface{}
	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is missing or not in JSON format"})
		return
	}

	if ghuToken == "" {
		ghuToken = strings.Split(c.GetHeader("Authorization"), " ")[1]
	}

	if ghuToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gho_token not found"})
		return
	}
	accToken, err := getAccToken(ghuToken)
	if accToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionId := fmt.Sprintf("%s%d", uuid.New().String(), time.Now().UnixNano()/int64(time.Millisecond))
	machineID := sha256.Sum256([]byte(uuid.New().String()))
	machineIDStr := hex.EncodeToString(machineID[:])
	accHeaders := getAccHeaders(accToken, uuid.New().String(), sessionId, machineIDStr)
	client := &http.Client{}

	jsonData, err := json.Marshal(jsonBody)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	isStream := gjson.GetBytes(jsonData, "stream").String() == "true"

	req, err := http.NewRequest("POST", completionsUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	for key, value := range accHeaders {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Printf("对话失败：%d, %s ", resp.StatusCode, bodyString)
		cache := cache.New(5*time.Minute, 10*time.Minute)
		cache.Delete(ghuToken)
		c.AbortWithError(resp.StatusCode, err)
		return
	}

	if isStream {
		returnStream(c, resp)
	} else {
		returnJson(c, resp)
	}

	c.Header("Content-Type", "text/event-stream; charset=utf-8")

	// 创建一个新的扫描器
	scanner := bufio.NewScanner(resp.Body)

	// 使用Scan方法来读取流
	for scanner.Scan() {
		line := scanner.Bytes()

		// 替换 "content":null 为 "content":""
		modifiedLine := bytes.Replace(line, []byte(`"content":null`), []byte(`"content":""`), -1)

		// 将修改后的数据写入响应体
		if _, err := c.Writer.Write(modifiedLine); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 添加一个换行符
		if _, err := c.Writer.Write([]byte("\n")); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	if scanner.Err() != nil {
		// 处理来自扫描器的任何错误
		c.AbortWithError(http.StatusInternalServerError, scanner.Err())
		return
	}
	return
}

func returnJson(c *gin.Context, resp *http.Response) {
	body, err := io.ReadAll(resp.Body.(io.Reader))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Writer.Write(body)
	return
}

func returnStream(c *gin.Context, resp *http.Response) {
	c.Header("Content-Type", "text/event-stream; charset=utf-8")

	// 创建一个新的扫描器
	scanner := bufio.NewScanner(resp.Body)

	// 使用Scan方法来读取流
	for scanner.Scan() {
		line := scanner.Bytes()

		// 替换 "content":null 为 "content":""
		modifiedLine := bytes.Replace(line, []byte(`"content":null`), []byte(`"content":""`), -1)

		// 将修改后的数据写入响应体
		if _, err := c.Writer.Write(modifiedLine); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 添加一个换行符
		if _, err := c.Writer.Write([]byte("\n")); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	if scanner.Err() != nil {
		// 处理来自扫描器的任何错误
		c.AbortWithError(http.StatusInternalServerError, scanner.Err())
		return
	}
	return
}

func models() ModelList {
	jsonStr := `{
        "object": "list",
        "data": [
            {"id": "text-search-babbage-doc-001","object": "model","created": 1651172509,"owned_by": "openai-dev"},
            {"id": "gpt-4-0613","object": "model","created": 1686588896,"owned_by": "openai"},
            {"id": "gpt-4", "object": "model", "created": 1687882411, "owned_by": "openai"},
            {"id": "babbage", "object": "model", "created": 1649358449, "owned_by": "openai"},
            {"id": "gpt-3.5-turbo-0613", "object": "model", "created": 1686587434, "owned_by": "openai"},
            {"id": "text-babbage-001", "object": "model", "created": 1649364043, "owned_by": "openai"},
            {"id": "gpt-3.5-turbo", "object": "model", "created": 1677610602, "owned_by": "openai"},
            {"id": "gpt-3.5-turbo-1106", "object": "model", "created": 1698959748, "owned_by": "system"},
            {"id": "curie-instruct-beta", "object": "model", "created": 1649364042, "owned_by": "openai"},
            {"id": "gpt-3.5-turbo-0301", "object": "model", "created": 1677649963, "owned_by": "openai"},
            {"id": "gpt-3.5-turbo-16k-0613", "object": "model", "created": 1685474247, "owned_by": "openai"},
            {"id": "text-embedding-ada-002", "object": "model", "created": 1671217299, "owned_by": "openai-internal"},
            {"id": "davinci-similarity", "object": "model", "created": 1651172509, "owned_by": "openai-dev"},
            {"id": "curie-similarity", "object": "model", "created": 1651172510, "owned_by": "openai-dev"},
            {"id": "babbage-search-document", "object": "model", "created": 1651172510, "owned_by": "openai-dev"},
            {"id": "curie-search-document", "object": "model", "created": 1651172508, "owned_by": "openai-dev"},
            {"id": "babbage-code-search-code", "object": "model", "created": 1651172509, "owned_by": "openai-dev"},
            {"id": "ada-code-search-text", "object": "model", "created": 1651172510, "owned_by": "openai-dev"},
            {"id": "text-search-curie-query-001", "object": "model", "created": 1651172509, "owned_by": "openai-dev"},
            {"id": "text-davinci-002", "object": "model", "created": 1649880484, "owned_by": "openai"},
            {"id": "ada", "object": "model", "created": 1649357491, "owned_by": "openai"},
            {"id": "text-ada-001", "object": "model", "created": 1649364042, "owned_by": "openai"},
            {"id": "ada-similarity", "object": "model", "created": 1651172507, "owned_by": "openai-dev"},
            {"id": "code-search-ada-code-001", "object": "model", "created": 1651172507, "owned_by": "openai-dev"},
            {"id": "text-similarity-ada-001", "object": "model", "created": 1651172505, "owned_by": "openai-dev"},
            {"id": "text-davinci-edit-001", "object": "model", "created": 1649809179, "owned_by": "openai"},
            {"id": "code-davinci-edit-001", "object": "model", "created": 1649880484, "owned_by": "openai"},
            {"id": "text-search-curie-doc-001", "object": "model", "created": 1651172509, "owned_by": "openai-dev"},
            {"id": "text-curie-001", "object": "model", "created": 1649364043, "owned_by": "openai"},
            {"id": "curie", "object": "model", "created": 1649359874, "owned_by": "openai"},
            {"id": "davinci", "object": "model", "created": 1649359874, "owned_by": "openai"},
            {"id": "gpt-4-0314", "object": "model", "created": 1687882410, "owned_by": "openai"}
        ]
    }`

	var modelList ModelList
	json.Unmarshal([]byte(jsonStr), &modelList)
	return modelList
}

func getAccToken(ghuToken string) (string, error) {
	var accToken = ""

	cache := cache.New(15*time.Minute, 60*time.Minute)
	cacheToken, found := cache.Get(ghuToken)
	if found {
		accToken = cacheToken.(string)
	} else {
		client := &http.Client{}
		req, err := http.NewRequest("GET", tokenUrl, nil)
		if err != nil {
			return accToken, err
		}

		headers := getHeaders(ghuToken)

		for key, value := range headers {
			req.Header.Add(key, value)
		}

		resp, err := client.Do(req)
		if err != nil {
			return accToken, err
		}
		defer resp.Body.Close()

		var reader interface{}
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(resp.Body)
			if err != nil {
				return accToken, fmt.Errorf("数据解压失败")
			}
		default:
			reader = resp.Body
		}

		body, err := io.ReadAll(reader.(io.Reader))
		if err != nil {
			return accToken, fmt.Errorf("数据读取失败")
		}
		if resp.StatusCode == http.StatusOK {
			accToken = gjson.GetBytes(body, "token").String()
			if accToken == "" {
				return accToken, fmt.Errorf("acc_token 未返回")
			}
			cache.Set(ghuToken, accToken, 14*time.Minute)
		} else {
			log.Printf("获取 acc_token 请求失败：%d, %s ", resp.StatusCode, string(body))
			return accToken, fmt.Errorf("获取 acc_token 请求失败： %d", resp.StatusCode)
		}
	}
	return accToken, nil
}

func getHeaders(ghoToken string) map[string]string {
	return map[string]string{
		"Host":          "api.github.com",
		"Authorization": "token " + ghoToken,

		"Editor-Version":        "vscode/1.85.0",
		"Editor-Plugin-Version": "copilot-chat/0.11.1",
		"User-Agent":            "GitHubCopilotChat/0.11.1",
		"Accept":                "*/*",
		"Accept-Encoding":       "gzip, deflate, br",
	}
}

func getAccHeaders(accessToken, uuid string, sessionId string, machineId string) map[string]string {
	return map[string]string{
		"Host":                   "api.githubcopilot.com",
		"Authorization":          "Bearer " + accessToken,
		"X-Request-Id":           uuid,
		"X-Github-Api-Version":   "2023-07-07",
		"Vscode-Sessionid":       sessionId,
		"Vscode-machineid":       machineId,
		"Editor-Version":         "vscode/1.85.0",
		"Editor-Plugin-Version":  "copilot-chat/0.11.1",
		"Openai-Organization":    "github-copilot",
		"Openai-Intent":          "conversation-panel",
		"Content-Type":           "application/json",
		"User-Agent":             "GitHubCopilotChat/0.11.1",
		"Copilot-Integration-Id": "vscode-chat",
		"Accept":                 "*/*",
		"Accept-Encoding":        "gzip, deflate, br",
	}
}
