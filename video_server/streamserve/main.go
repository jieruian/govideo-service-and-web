package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMidlleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "请求太多了")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegistHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	router.GET("/testpages", testPageHandler)
	return router
}

func main() {
	r := RegistHandler()
	mh := NewMidlleWareHandler(r, 1)
	http.ListenAndServe("127.0.0.1:9000", mh)
}
