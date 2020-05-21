package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegistHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	return router
}

func main() {
  r := RegistHandler()
  http.ListenAndServe("127.0.0.1:9000".r)
}
