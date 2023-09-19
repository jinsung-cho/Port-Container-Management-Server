package controller

import (
	"fmt"
	"net/http"
)

func GetCheckpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetCheckpoint")
}

func GetCheckpointState(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetCheckpointState")
}