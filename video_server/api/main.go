package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statue := validateUserSession(r)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	r.Header.Set("Access-Control-Allow-Origin", "*")

	fmt.Println(r.Header.Get("Content-Type"))
	if statue {
		fmt.Println("不需要重新登录")
	} else {
		fmt.Println("需要重登陆")
	}

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	//注册
	router.POST("/user", RegistUser)
	//登录
	router.POST("/user/name=:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe("127.0.0.1:8009", mh)
	fmt.Println("-----")
}
