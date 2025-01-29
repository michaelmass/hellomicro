// Taken from https://github.com/grpc-ecosystem/grpc-gateway/issues/4263#issuecomment-2081459161

package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

type etagWriter struct {
	http.ResponseWriter
	wroteHeader bool
	ifNoneMatch string
	writeEtag   bool
	minBytes    int
}

func (w etagWriter) Write(b []byte) (int, error) {
	if w.wroteHeader || !w.writeEtag || len(b) < w.minBytes {
		return w.ResponseWriter.Write(b)
	}

	// Generate the Etag
	h := md5.New()
	h.Write(b)
	etag := fmt.Sprintf("\"%s\"", hex.EncodeToString(h.Sum(nil)))

	w.Header().Set("Etag", etag)

	if w.ifNoneMatch != "" && w.ifNoneMatch == etag {
		w.ResponseWriter.WriteHeader(http.StatusNotModified)
		return 0, nil
	} else {
		return w.ResponseWriter.Write(b)
	}
}

func (w etagWriter) WriteHeader(code int) {
	// Track if the headers have already been written
	w.ResponseWriter.WriteHeader(code)
}

// Unwrap returns the original http.ResponseWriter. This is necessary
// to expose Flush() and Push() on the underlying response writer.
func (w etagWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

// EtagHandler wraps an http.Handler and will write an Etag header to the
// response if the request method is GET and the response size is greater
// than or equal to minBytes.  It will also return a NotModified response
// if the If-None-Match header matches the Etag header.
func EtagHandler(h http.Handler, minBytes int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w = etagWriter{
			ResponseWriter: w,
			ifNoneMatch:    r.Header.Get("If-None-Match"),
			writeEtag:      r.Method == http.MethodGet,
			minBytes:       minBytes,
		}

		h.ServeHTTP(w, r)
	})
}
