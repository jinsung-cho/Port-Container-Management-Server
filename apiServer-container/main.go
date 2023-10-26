package main

import (
	"net/http"
	"os"
	"server/router"
	"server/util"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")
	util.CheckRuntimeError(env_err, ".env Load fail")

	err := http.ListenAndServe(":"+os.Getenv("API_SERVER_PORT"), router.InitRouter())
	util.CheckRuntimeError(err, "http open err")
}
