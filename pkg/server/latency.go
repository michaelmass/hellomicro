package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func latency(w http.ResponseWriter, req *http.Request) {
	duration, err := strconv.Atoi(getParam(req, "duration"))

	if err != nil {
		badRequest(w, err)
		return
	}

	time.Sleep(time.Duration(duration) * time.Millisecond)
	w.Write([]byte(fmt.Sprintf("waited %d ms", duration)))
}
