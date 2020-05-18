package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func RegisterHandlers()  *httprouter.Router{
	router := httprouter.New()
    //注册
	router.POST("/user",RegistUser)
	//登录
	router.POST("/user/name=:user_name",Login)


	return router
}

func main() {
  r := RegisterHandlers()
  http.ListenAndServe("127.0.0.1:8009",r)
  fmt.Println("-----")
}


func testFunc(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w,"这是创建用户啦😜😜")
}