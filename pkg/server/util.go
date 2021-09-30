package server

import (
	"encoding/json"
	"net/http"
)

func getParam(req *http.Request, name string) string {
	return req.URL.Query().Get(name)
}

func writeJson(w http.ResponseWriter, data interface{}) {
	content, err := json.Marshal(data)

	if err != nil {
		badRequest(w, err)
		return
	}

	w.Write(content)
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
