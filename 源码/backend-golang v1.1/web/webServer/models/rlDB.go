package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Regist struct { // 注册用的结构体
	Birthday     string `json:"birthday"`
	Gender       string `json:"gender"`       // 性别
	Introduction string `json:"introduction"` // 简介
	Password     string `json:"password"`
	PhoneNumber  string `json:"phoneNumber"`
	Portrait     string `json:"portrait"` // 头像
	UserName     string `json:"userName"`
}

func IsTelephoneExists(PhoneNumber string) (int, bool) { // 查找手机号是否存在
	userID := 0
	//具体的以这个为准
	sqlStr := `select userAccount
	from userinfo
	where phoneNumber = ?`
	//下面这个只是测试用的
	// sqlStr := `select uA
	// from userinfo
	// where phoneNumber = ?`
	err := db.QueryRow(sqlStr, PhoneNumber).Scan(&userID)
	fmt.Println("查找手机号错误：", err)
	return userID, err == nil
}

func CreateUser(registInfo Regist, registTime string) bool { // 创建用户
	sqlStr := `insert into userinfo(userName, password, gender,portrait,introduction,birthday,phoneNumber,registTime) values(?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, registInfo.UserName, registInfo.Password, registInfo.Gender, registInfo.Portrait, registInfo.Introduction, registInfo.Birthday, registInfo.PhoneNumber, registTime)
	fmt.Println(err)
	return err == nil
}

func SecretCorrect(PhoneNumber string, password string) bool { // 判断密码是否正确
	sqlStr := `select password
	from userinfo 
	where phoneNumber = ?`
	var pwd string
	err := db.QueryRow(sqlStr, PhoneNumber).Scan(&pwd)
	if err != nil {
		return false // 找不到密码
	}
	if pwd != password {
		return false // 密码错误
	}
	return true
}

func SelectAll(phone string) Regist {
	var rg Regist
	sqlStr := `select userName, password, gender,portrait,introduction,birthday,phoneNumber
	from userinfo 
	where phoneNumber = ?`
	err := db.QueryRow(sqlStr, phone).Scan(&rg.UserName, &rg.Password, &rg.Gender, &rg.Portrait, &rg.Introduction, &rg.Birthday, &rg.PhoneNumber)
	if err != nil {
		return rg
	}
	return rg
}
