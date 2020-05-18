package dbops

import (
	"database/sql"
	"fmt"
	"govideo/video_server/api/defs"
	"govideo/video_server/api/utils"
	"time"
)

//<editor-fold desc="users表的操作">
//注册用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := db.Prepare("insert into users (login_name,pwd) value (?,?)")
	if err != nil {
		fmt.Println("注册数据库错误", err)
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

//获取用户
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := db.Prepare("select pwd from users where login_name=?")
	if err != nil {
		fmt.Println("获取用户数据库错误", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := db.Prepare("delete from users where login_name=? and pwd=?")
	if err != nil {
		fmt.Println("删除用户数据库错误", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//</editor-fold>

//<editor-fold desc="video_info表的操作">  op+cmd+T
//增加新视频  aid 作者的id    name 视频的名字
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := db.Prepare(`insert into video_info (id, author_id, name, display_ctime) value (?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}
	defer stmtIns.Close()
	return res, nil
}

//根据videoId获取video
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := db.Prepare("select author_id, name, display_ctime from video_info where id=?")
	var (
		aid  int
		name string
		dct  string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	return &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: dct,
	}, nil
}
func DeleteVideoInfo(vid string) error {
	stmtDel, err := db.Prepare("delete from video_info where id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//</editor-fold>
