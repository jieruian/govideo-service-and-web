package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"govideo/video_server/api/session"
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
	//获取用户信息
	router.GET("/user/:username", CreateUserInfo)
	//上传新的video
	router.POST("/user/:username/psotvideos", AddNewVideo)
	//视频列表
	router.GET("/user/:username/getvideos", ListAllVideos)
	//删除视频
	router.DELETE("/user/:username/deletevideos/:vid-id", DeleteVideo)
	//提交评论
	router.POST("/videos/:vid-id/postcomments", PostComment)
	//显示评论
	router.GET("/videos/:vid-id/getcomments", ShowComments)

	return router
}

func main() {
	prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe("127.0.0.1:8009", mh)
	fmt.Println("-----")
}

func prepare() {
	session.LoadSessionFromDB()
}
