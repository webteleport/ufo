package echo

import (
	"expvar"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	echo "github.com/jpillora/go-echo-server/handler"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

func getVar(s string) (result *expvar.String) {
	expvar.Do(func(kv expvar.KeyValue) {
		if kv.Key == s {
			result = kv.Value.(*expvar.String)
		}
	})
	if result != nil {
		return
	}
	return expvar.NewString(s)
}

func Run(args []string) error {
	e := echo.New(echo.Config{})
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			e.ServeHTTP(w, r)
			return
		}
		lastSegment := path.Base(path.Clean(r.URL.Path))
		bodyVar := getVar(lastSegment)
		bodyStr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()
		switch r.Method {
		case http.MethodPut, http.MethodPost:
			bodyVar.Set(string(bodyStr))
		case http.MethodDelete:
			bodyVar.Set("")
		}
		io.WriteString(w, bodyVar.Value())
	})
	mux.Handle("/debug/vars", expvar.Handler())
	return wtf.Serve(Arg0(args, "https://ufo.k0s.io"), utils.GinLoggerMiddleware(mux))
}
