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
	//defer db.Close()

	http.ListenAndServe(":"+os.Getenv("HOST_PORT"), router.InitRouter())
}
