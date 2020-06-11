package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func registHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userhomeHandler)
	router.POST("/userhome", userhomeHandler)
	router.POST("/api", apiHandler)
	router.POST("/upload/:vid-id", proxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	return router
}

func main() {
	r := registHandler()
	http.ListenAndServe("127.0.0.1:9003", r)
}
