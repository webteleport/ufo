package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/webteleport/utils"
)

func Handler(s string) http.Handler {
	handler := http.FileServer(http.Dir(s))
	handler = utils.GzipMiddleware(handler)
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
		fmt.Println("Saving", dst.Name(), v.Size)
		uploadedFile, _ := v.Open()
		dst, _ := os.Create(v.Filename)
		io.Copy(dst, uploadedFile)
		uploadedFile.Close()
		dst.Close()
	}

	io.WriteString(w, "Upload succeeded.")
}
