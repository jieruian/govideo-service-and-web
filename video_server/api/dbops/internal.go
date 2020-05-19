package dbops

import (
	"database/sql"
	"govideo/video_server/api/defs"
	"strconv"
	"sync"
)

//插入session
func InsertSession(sid string, ttl int64, umane string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := db.Prepare("insert sessions (session_id, TTL, login_name) value (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlStr, umane)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

//根据sessionid获取session对象
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := db.Prepare("select TTL, login_name from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

//取出所有的session
func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := db.Prepare("select * from sessions ")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
LOOP:
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			break LOOP
		}
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			//log.Printf("session id: 5s, ttl: %d", id, ss.TTL)
		}
	}
	defer stmtOut.Close()
	return m, nil
}

//删除session
func DeleteSession(sid string) error {
	stmtOut, err := db.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		return err
	}
	if _, err := stmtOut.Exec(sid); err != nil {
		return err
	}
	defer stmtOut.Close()
	return nil
}
