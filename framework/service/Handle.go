package service

import (
	"net/http"
	"mime"
	"path/filepath"
	"os"
	"io"
)

func HandleStatic(path string, dir string) {
	pl := len(path)
	http.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		ctype := mime.TypeByExtension(filepath.Ext(req.URL.Path[pl:]))
		src, err := os.Open(dir + "/" + req.URL.Path[pl:])
		defer src.Close()
		if err != nil {
			res.WriteHeader(404)
			return
		}
		res.Header().Set("Content-Type", ctype)
		io.Copy(res, src)
	})
}