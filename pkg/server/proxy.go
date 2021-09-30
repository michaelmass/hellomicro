package server

import (
	"io"
	"net/http"
	"strings"
)

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func proxy(w http.ResponseWriter, req *http.Request) {
	method := getParam(req, "method")
	url := getParam(req, "url")
	body := getParam(req, "body")

	request, err := http.NewRequest(method, url, strings.NewReader(body))

	if err != nil {
		badRequest(w, err)
		return
	}

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		badRequest(w, err)
		return
	}

	copyHeader(w.Header(), res.Header)
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}
