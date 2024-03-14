package handler

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/webteleport/utils"
)

//go:embed index.html
var WWW string

var WWWHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, WWW)
})

func Handler(s string) http.Handler {
	handler := utils.GzipMiddleware(WWWHandler)
	handler = utils.GinLoggerMiddleware(handler)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	mux.HandleFunc("POST /", upload)
	return mux
}

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(8 * (1 << 30)) // set max file size to 8GB
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	for _, v := range r.MultipartForm.File["files"] {
		uploadedFile, _ := v.Open()
		dst, _ := os.Create(v.Filename)
		fmt.Println("Saving", dst.Name(), v.Size)
		io.Copy(dst, uploadedFile)
		uploadedFile.Close()
		dst.Close()
	}

	io.WriteString(w, "Upload succeeded.")
}
