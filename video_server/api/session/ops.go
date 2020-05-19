package session

import (
	"govideo/video_server/api/dbops"
	"govideo/video_server/api/defs"
	"govideo/video_server/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func noInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

//从数据库里删除过期session
func DeleteExpiredSession(id string) {
	dbops.DeleteSession(id)
	sessionMap.Delete(id)
}

//从数据库获取session
func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		//ss := v
		sessionMap.Store(k, ss)
		return true
	})
}

//根据用户名产生一个新的session
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := noInMilli()
	ttl := ct + 30*60*1000 // 过期时间为30分钟 Server side session valid time: 30 min
	dbops.InsertSession(id, ttl, un)
	ss := &defs.SimpleSession{
		Username: un,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	return id
}

//根据sessionId判断session过期，如果没有过期返回userName，false。过期返回null，true
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := noInMilli()
		if ct > ss.(*defs.SimpleSession).TTL {
			//过期了
			DeleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		return "", true
	}
}
