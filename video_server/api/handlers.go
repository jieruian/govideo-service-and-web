package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func RegistUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
    io.WriteString(w,"这是创建用户")
}

func Login(w http.ResponseWriter, r *http.Request,p httprouter.Params)  {
    userName := p.ByName("user_name")
    io.WriteString(w,userName)
}