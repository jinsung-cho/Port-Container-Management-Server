package main

import(
	"net/http"
	"server/router"
)


func main(){
	routeHandler := router.InitRouter()
	http.ListenAndServe(":8000", routeHandler)
}