package main

import (
	"net/http"
	"os"
	"server/router"
	"server/util"

	"github.com/joho/godotenv"
)

func main(){
	env_err := godotenv.Load(".env")
	util.CheckRuntimeError(env_err, ".env Load fail")
	
	// routeHandler := router.InitRouter()
	http.ListenAndServe(":" + os.Getenv("HOST_PORT"), router.InitRouter())
}