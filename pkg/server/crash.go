package server

import (
	"net/http"
	"os"
)

func exit(w http.ResponseWriter, req *http.Request) {
	os.Exit(1)
}

func crash(w http.ResponseWriter, req *http.Request) {
	go func() { panic("oops") }()
}
