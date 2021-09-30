package server

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func root(handlers map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")

		html := "<h1>Hello Micro</h1>"

		for path := range handlers {
			html += fmt.Sprintf("<div><a href=\"%s\"><h3>%s</h3></a></div>", path, path)
		}

		w.Write([]byte(html))
	}
}

func Run() {
	handlers := map[string]http.HandlerFunc{
		"/ping":       ping,
		"/echo":       echo,
		"/info":       info,
		"/env":        env,
		"/crash":      crash,
		"/exit":       exit,
		"/latency":    latency,
		"/log":        log,
		"/proxy":      proxy,
		"/files/list": listFiles,
		"/files/put":  putFile,
		"/metrics":    promhttp.Handler().ServeHTTP,
	}

	http.HandleFunc("/", root(handlers))

	for path, handler := range handlers {
		http.HandleFunc(path, handler)
	}

	http.ListenAndServe(":8080", nil)
}
