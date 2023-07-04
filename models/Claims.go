package models

type LoginInfo struct {
	Password string `json:"password"` // 密码
	UserID   int64  `json:"userId"`   // 用户编号
	Username string `json:"userName"` // 用户名
}

type User struct { //登录验证
	UserName  string
	UserClaim int //用户声明，
}

// type Claims struct {
// 	ID     int    `json:"claim_id"`
// 	AuthID int    `json:"auth_id"`
// 	Type   string `json:"type"`
// 	Value  string `json:"value"`
// }

// func GetUserClaims(userName string) (claim int) {
// 	var auth LoginInfo
// 	sqlstr := "SELECT userAccount FROM userInfo WHERE userName=?"
// 	err := db.QueryRow(sqlstr, userName).Scan(&auth.UserID)
// 	if err != nil {
// 		fmt.Printf("query failed, err:%v\n", err)
// 		return -1
// 	}
// 	if auth.UserID > 0 {
// 		// fmt.Print(buser.UserID)
// 		return true
// 	}
// 	return false
// 	// db.Where("username = ?", userName).First(&auth)
// 	// db.Where("auth_id = ?", auth.ID).Find(&claims)
// 	return
// }
