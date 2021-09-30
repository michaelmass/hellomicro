package server

import (
	"net/http"
	"os"
	"strings"
)

func env(w http.ResponseWriter, req *http.Request) {
	environments := map[string]string{}

	for _, line := range os.Environ() {
		keyval := strings.Split(line, "=")
		environments[keyval[0]] = keyval[1]
	}

	writeJson(w, environments)
}
