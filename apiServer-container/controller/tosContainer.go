package controller

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"server/util"
)

func GetTosContainer(w http.ResponseWriter, r *http.Request) {
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/container/tos"

	bodyBytes, err := io.ReadAll(r.Body)
	if util.CheckHttpError(w, err, "Reading body") {
		return
	}
	bodyReader := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest("POST", url, bodyReader)

	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if util.CheckHttpError(w, err, "Check Client Do") {
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if util.CheckHttpError(w, err, "Check Copying Response") {
		return
	}
}
