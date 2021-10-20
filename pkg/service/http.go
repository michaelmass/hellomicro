package service

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func newSpa(staticPath, indexPath string) *spaHandler {
	return &spaHandler{
		staticPath: staticPath,
		indexPath:  indexPath,
	}
}

func (service *Service) handler(h http.Handler) http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/v1").Handler(h)
	r.PathPrefix("/").Handler(newSpa("static", "index.html"))

	return r
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
