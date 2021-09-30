package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/dustin/go-humanize"
)

func listFiles(w http.ResponseWriter, req *http.Request) {
	folder := getParam(req, "folder")

	var filesInfo []map[string]string

	files, err := filepath.Glob(folder + "/*")

	if err != nil {
		badRequest(w, err)
		return
	}

	for _, file := range files {
		fileInfo, err := os.Stat(file)

		if err != nil {
			badRequest(w, err)
			return
		}

		filesInfo = append(filesInfo, map[string]string{
			"path":    file,
			"size":    humanize.Bytes(uint64(fileInfo.Size())),
			"modTime": humanize.Time(fileInfo.ModTime()),
			"mode":    fileInfo.Mode().String(),
		})
	}

	writeJson(w, filesInfo)
}

func putFile(w http.ResponseWriter, req *http.Request) {
	filepath := getParam(req, "filepath")
	filename := getParam(req, "filename")

	err := ioutil.WriteFile(path.Join(filepath, filename), []byte(""), 0666)

	if err == nil {
		badRequest(w, err)
	}
}
