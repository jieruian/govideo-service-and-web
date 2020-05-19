package session

import "sync"

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

//从数据库获取session
func LoadSessionFromDB() {

}

//根据用户名产生一个新的session
func GenerateNewSessionId(un string) string {

}

//根据sessionId判断session过期，如果没有过期返回userName，false。过期返回null，true
func IsSessionExpired(sid string) (string, bool) {

}
