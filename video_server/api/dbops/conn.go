package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db  *sql.DB
	err error
)

func init() {
	dsn := "root:@tcp(localhost:3306)/video_server?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("打开数据库失败", err)
		//return err
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("校验失败", err)
		panic(err.Error())
	}
	fmt.Println("打开成功")
	log.Println("数据库准备开始啦")
}
