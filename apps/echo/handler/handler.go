package handler

import (
	"expvar"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	echo "github.com/jpillora/go-echo-server/handler"
)

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

func Handler() http.Handler {
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
	return mux
}
