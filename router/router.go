package router

import (
	"net/http"
	"server/controller"
	"server/util"
)

func InitRouter() http.Handler {
	env_err := godotenv.Load("../.env")
	util.FailOnError(env_err, ".env Load fail")
	hostIP := os.Getenv("HOST_IP")

	r := mux.NewRouter()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://" + hostIP + ":3000", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	handler := corsConfig.Handler(r)

	r.HandleFunc("/container", ContainerRouter).Methods("GET")
	r.HandleFunc("/container/spec", ContainerRouter).Methods("GET")
	r.HandleFunc("/container/spec/review", ContainerRouter).Methods("POST")
	r.HandleFunc("/checkpoint", SecurityCheckpointRouter).Methods("GET")
	r.HandleFunc("/checkpoint/state", SecurityCheckpointRouter).Methods("GET")

	return handler
}

func ContainerRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	}
}

func SecurityCheckpointRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	}
}