package controller

import (
	"fmt"
	"net/http"
)

func GetContainer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetContainer")
}

func GetContainerSpec(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetContainerSpec")
}

func AppendContainerReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AppendContainerReview")
}