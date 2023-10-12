package controller

import (
	"io"
	"net/http"
	"os"
	"server/util"
)

func GetContainer(w http.ResponseWriter, r *http.Request) {
	cntrNoHeader := r.Header.Get("cntrNo")
	if cntrNoHeader == "" {
		http.Error(w, "cntrNo and truckNo headers are required", http.StatusBadRequest)
		return
	}
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/container"

	req, err := http.NewRequest("GET", url, nil)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	req.Header.Set("cntrNo", cntrNoHeader)

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

func GetAllContainers(w http.ResponseWriter, r *http.Request) {
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/containers"

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

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/container"

	req, err := http.NewRequest("POST", url, r.Body)
	if util.CheckHttpError(w, err, "Check NewRequest") {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if util.CheckHttpError(w, err, "Check Client Do") {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to save data on remote server"))
	}
}

func GetAllContainerSpec(w http.ResponseWriter, r *http.Request) {
	dbServerHost := os.Getenv("DB_SERVER_HOST")
	dbServerPort := os.Getenv("DB_SERVER_PORT")
	url := "http://" + dbServerHost + ":" + dbServerPort + "/container/spec"

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
