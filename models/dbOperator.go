package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 全局对象db
var db *sql.DB

// 初始化数据库
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/webData"
	// 不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	defer db.Close()
}
