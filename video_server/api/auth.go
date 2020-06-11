package main

import (
	"govideo/video_server/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

//校验session
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		//请求头里没有sessionid
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		//session 过期了
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true

}

//校验这个用户
func validateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		//sendErrorResponse(w, )
		return false
	}
	return true
}
