package main

import (
	"db-server/model"
	"db-server/router"
	"db-server/util"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")
	util.CheckRuntimeError(env_err, ".env Load fail")

	model.InitDB()
	err := http.ListenAndServe(":"+os.Getenv("DB_SERVER_PORT"), router.InitRouter())
	util.CheckRuntimeError(err, "http open err")
}
