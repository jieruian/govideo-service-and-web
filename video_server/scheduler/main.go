package main

import (
	"github.com/julienschmidt/httprouter"
	"govideo/video_server/scheduler/taskrunner"
	"net/http"
)

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	r := registerHandlers()
	go taskrunner.Start()
	http.ListenAndServe("127.0.0.1:9001", r)
}
