package server

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func log(w http.ResponseWriter, req *http.Request) {
	count, err := strconv.Atoi(getParam(req, "count"))

	if err != nil {
		badRequest(w, err)
		return
	}

	uuid, err := uuid.NewRandom()

	if err != nil {
		badRequest(w, err)
		return
	}

	for i := 1; i <= count; i++ {
		logrus.Infof("%s: %d", uuid.String(), i)
	}
}
