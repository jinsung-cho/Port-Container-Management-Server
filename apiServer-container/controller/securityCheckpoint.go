package controller

import (
	"io"
	"net/http"
	"os"
	"server/util"
)

func GetAllCheckpoint(w http.ResponseWriter, r *http.Request) {
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/checkpoint"

	req, err := http.NewRequest("GET", url, nil)
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

func CreateCheckpoint(w http.ResponseWriter, r *http.Request) {

}

func GetAllCheckpointState(w http.ResponseWriter, r *http.Request) {
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/checkpoint/state"

	req, err := http.NewRequest("GET", url, nil)
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

func CreateCheckpointState(w http.ResponseWriter, r *http.Request) {

}
