package server

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestEcho struct {
	URL           *url.URL
	Method        string
	Proto         string
	Host          string
	ContentLength int64
	Header        http.Header
	Cookies       []*http.Cookie
	RemoteAddr    string
	RequestURI    string
	SourceIp      string
	Params        map[string][]string
	UserAgent     string
	Body          string
}

func formatRequest(r *http.Request) (*RequestEcho, error) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	return &RequestEcho{
		URL:           r.URL,
		Params:        r.URL.Query(),
		Method:        r.Method,
		Header:        r.Header,
		Proto:         r.Proto,
		ContentLength: r.ContentLength,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		UserAgent:     r.UserAgent(),
		Cookies:       r.Cookies(),
		Body:          string(body),
	}, nil
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong\n"))
}

func echo(w http.ResponseWriter, req *http.Request) {
	request, err := formatRequest(req)

	if err != nil {
		badRequest(w, err)
		return
	}

	writeJson(w, request)
}
