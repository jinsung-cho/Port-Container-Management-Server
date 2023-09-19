package router

import (
	"db-server/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRouter() http.Handler {
	hostIP := os.Getenv("HOST_IP")

	r := mux.NewRouter()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://" + hostIP + ":3000", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	// Container routes
	r.HandleFunc("/container", controller.GetContainer).Methods("GET")
	r.HandleFunc("/container", controller.GetContainer).Methods("POST")
	r.HandleFunc("/container/spec", controller.GetContainerSpec).Methods("GET")
	r.HandleFunc("/container/spec/review", controller.GetContainerRevuew).Methods("GET")
	r.HandleFunc("/container/spec/review", controller.AppendContainerReview).Methods("POST")

	// Security checkpoint routes
	r.HandleFunc("/checkpoint", controller.GetCheckpoint).Methods("GET")
	r.HandleFunc("/checkpoint/state", controller.GetCheckpointState).Methods("GET")

	handler := corsConfig.Handler(r)

	return handler
}
