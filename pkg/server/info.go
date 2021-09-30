package server

import (
	"net/http"
	"runtime"

	"github.com/dustin/go-humanize"
)

func info(w http.ResponseWriter, req *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	data := map[string]interface{}{
		"OS":           runtime.GOOS,
		"Version":      runtime.Version(),
		"NumCPU":       runtime.NumCPU(),
		"NumGoRoutine": runtime.NumGoroutine(),
		"NumCgoCall":   runtime.NumCgoCall(),
		"memory": map[string]interface{}{
			"Alloc":      humanize.Bytes(m.Alloc),
			"TotalAlloc": humanize.Bytes(m.TotalAlloc),
			"Sys":        humanize.Bytes(m.Sys),
			"NumGC":      m.NumGC,
		},
	}

	writeJson(w, data)
}
