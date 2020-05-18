package dbops

import (
	"fmt"
)
//注册用户
func AddUserCredential(loginName string,pwd string) error  {
	stmtIns,err := db.Prepare("insert into users (login_name,pwd) value (?,?)")
	if err != nil {
	   fmt.Println("注册数据库错误",err)
	   return err
	}
    stmtIns.Exec(loginName,pwd)
	stmtIns.Close()
	return nil
}
//获取用户
func GetUserCredential(loginName string)(string, error) {
  stmtOut,err := db.Prepare("select pwd from users where login_name")
  if err != nil {
	  fmt.Println("获取用户数据库错误",err)
	  return "",err
  }
  var pwd string
  stmtOut.QueryRow(loginName).Scan(&pwd)
  stmtOut.Close()
  return pwd,nil
}

func DeleteUser(loginName string,pwd string) error {
	stmtDel,err := db.Prepare("delete from users where login_name=? and pwd=?")
	if err != nil {
		fmt.Println("删除用户数据库错误",err)
		return err
	}
	stmtDel.Exec(loginName,pwd)
	stmtDel.Close()
	return nil
}


