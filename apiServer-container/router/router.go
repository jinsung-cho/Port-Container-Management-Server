package router

import (
	"net/http"
	"os"
	"server/controller"

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
	r.HandleFunc("/containers", controller.GetAllContainers).Methods("GET")
	r.HandleFunc("/container", controller.CreateContainer).Methods("POST")
	r.HandleFunc("/container/spec", controller.GetAllContainerSpec).Methods("GET")
	r.HandleFunc("/container/spec/review", controller.GetContainerReview).Methods("GET")
	r.HandleFunc("/container/spec/review", controller.AppendContainerReview).Methods("POST")

	// Security checkpoint routes
	r.HandleFunc("/checkpoint", controller.GetAllCheckpoint).Methods("GET")
	r.HandleFunc("/checkpoint", controller.CreateCheckpoint).Methods("POST")
	r.HandleFunc("/checkpoint/state", controller.GetAllCheckpointState).Methods("GET")
	r.HandleFunc("/checkpoint/state", controller.CreateCheckpointState).Methods("POST")

	handler := corsConfig.Handler(r)

	return handler
}
