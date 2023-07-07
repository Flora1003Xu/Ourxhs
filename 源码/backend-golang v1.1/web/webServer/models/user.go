package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 用户基本信息
type UserInfo struct {
	Birthday     string `json:"birthday" form:"birthday"`
	CollectedNum int64  `json:"collectedNum" form:"collectedNum"` // 被收藏的笔记数
	CollectNum   int64  `json:"collectNum" form:"collectNum"`     // 收藏数
	FansNum      int64  `json:"fansNum" form:"fansNum"`
	FollowNum    int64  `json:"followNum" form:"followNum"`       // 关注数
	Gender       string `json:"gender" form:"gender"`             // 性别
	Introduction string `json:"introduction" form:"introduction"` // 简介
	LikedNum     int64  `json:"likedNum" form:"likedNum"`         // 被点赞数量
	NoteNum      int64  `json:"noteNum" form:"noteNum"`
	Password     string `json:"password" form:"password"`
	PhoneNumber  string `json:"phoneNumber" form:"phoneNumber"`
	Portrait     string `json:"portrait" form:"portrait"` // 头像
	RegistTime   string `json:"registTime" form:"registTime"`
	UserID       int64  `json:"userAccount" form:"userAccount"`
	UserName     string `json:"userName" form:"userName"`
}

// 用户可修改信息
type ModifiableInfo struct {
	Birthday     string `json:"birthday" form:"birthday"`
	Gender       string `json:"gender" form:"gender"`             // 性别
	Introduction string `json:"introduction" form:"introduction"` // 简介
	Password     string `json:"password" form:"password"`
	Portrait     string `json:"portrait" form:"portrait"` // 头像
	UserName     string `json:"userName" form:"userName"`
}

func UserInfoDB(id int) UserInfo { // 从数据库获得用户信息
	sqlStr := `select userAccount, userName, gender, portrait, introduction, fansNum, noteNum, collectNum, followNum, collectedNum, likedNum, phoneNumber, password,birthday,registTime
	from userinfo 
	where userAccount = ?`
	var ui UserInfo
	var bd, rt string
	err := db.QueryRow(sqlStr, id).Scan(&ui.UserID, &ui.UserName, &ui.Gender, &ui.Portrait, &ui.Introduction, &ui.FansNum, &ui.NoteNum, &ui.CollectNum, &ui.FollowNum, &ui.CollectedNum, &ui.LikedNum, &ui.PhoneNumber, &ui.Password, &bd, &rt)
	ui.Birthday = bd
	ui.RegistTime = rt
	// 如果要使用time.Time()类型的birthday/registTime，则将上面两句改成下面两句
	// ui.Birthday, _ = time.Parse("2006-01-02 15:04:05", bd)
	// ui.RegistTime, _ = time.Parse("2006-01-02 15:04:05", rt)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ui
	}
	fmt.Println(ui.UserID, ui.UserName, ui.Gender, ui.Portrait, ui.Introduction, ui.Birthday, ui.RegistTime, ui.FansNum, ui.NoteNum, ui.CollectNum, ui.FollowNum, ui.CollectedNum, ui.LikedNum, ui.PhoneNumber, ui.Password)
	return ui
}

// 从数据库获得某用户发布的笔记信息
func NoteInfoDB(id int) []Note {
	var notes []Note
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.userName, u.portrait
	from noteInfo n, userInfo u
	where creatorAccount = ? and n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return notes
	}
	defer rows.Close()
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.CreatorName, &nt.Portrait)
		fmt.Println(nt.NoteID, nt.Title, nt.Cover, nt.CreatorID, nt.LikedNum, nt.CreatorName, nt.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return notes
		}
		notes = append(notes, nt)
	}
	return notes
}

// 从数据库获得某用户收藏的笔记信息
func CollectInfoDB(id int) []Note {
	var collects []Note
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.userName, u.portrait
	from noteInfo n, userInfo u, collectTable c
	where c.userAct = ? and c.collectNoteId=n.noteId and n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return collects
	}
	defer rows.Close()
	for rows.Next() {
		var ct Note
		err := rows.Scan(&ct.NoteID, &ct.Title, &ct.Cover, &ct.CreatorID, &ct.LikedNum, &ct.CreatorName, &ct.Portrait)
		fmt.Println(ct.NoteID, ct.Title, ct.Cover, ct.CreatorID, ct.LikedNum, ct.CreatorName, ct.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return collects
		}
		collects = append(collects, ct)
	}
	return collects
}

// 从数据库获得某用户点赞的笔记信息
func LikeInfoDB(id int) []Note {
	var collects []Note
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.userName, u.portrait
	from noteInfo n, userInfo u, favorTable c
	where c.userAct = ? and c.favorNoteId=n.noteId and n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return collects
	}
	defer rows.Close()
	for rows.Next() {
		var ct Note
		err := rows.Scan(&ct.NoteID, &ct.Title, &ct.Cover, &ct.CreatorID, &ct.LikedNum, &ct.CreatorName, &ct.Portrait)
		fmt.Println(ct.NoteID, ct.Title, ct.Cover, ct.CreatorID, ct.LikedNum, ct.CreatorName, ct.Portrait)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return collects
		}
		collects = append(collects, ct)
	}
	return collects
}

// 修改用户的信息并存放到数据库中
func ModifyInfo(beforeInfo ModifiableInfo, id int, yes bool) bool {
	var result sql.Result
	var err error
	if yes {
		sqlStr := `update userInfo set birthday = ?, gender = ?, introduction = ?, portrait = ?, userName = ? 
		where userAccount = ?`
		//birTime, _ := time.ParseInLocation("2006-01-02", beforeInfo.Birthday, time.Local) // string转time
		result, err = db.Exec(sqlStr, beforeInfo.Birthday, beforeInfo.Gender, beforeInfo.Introduction, beforeInfo.Portrait, beforeInfo.UserName, id)
	} else {
		sqlStr := `update userInfo set birthday = ?, gender = ?, introduction = ?, userName = ? 
		where userAccount = ?`
		//birTime, _ := time.ParseInLocation("2006-01-02", beforeInfo.Birthday, time.Local) // string转time
		result, err = db.Exec(sqlStr, beforeInfo.Birthday, beforeInfo.Gender, beforeInfo.Introduction, beforeInfo.UserName, id)
	}
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	idAff, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed:", err)
		return false
	}
	fmt.Println("id:", idAff)
	return true
}

// 修改用户发布笔记数
func ChangeNoteNum(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set noteNum =noteNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set noteNum =noteNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("用户编号：%d\n", n)
	// return
}

// 修改用户获赞数
func ChangeUserLikes(noteId, option int) {
	var sqlstr string
	userId := NoteToUser(noteId)
	addnum := `UPDATE userInfo set likedNum =likedNum+1 WHERE userAccount = ? `
	reducenum := `UPDATE userInfo set likedNum =likedNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("用户编号：%d\n", n)
}

// 修改用户被收藏笔记数量
func ChangeUserCollects(noteId, option int) {
	var sqlstr string
	userId := NoteToUser(noteId)
	addnum := `UPDATE userInfo set collectedNum =collectedNum+1 WHERE userAccount = ? `
	reducenum := `UPDATE userInfo set collectedNum =collectedNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("用户被收藏数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("用户被收藏get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("收藏数增加的用户编号：%d\n", n)
}

// 获取用户关注列表
func GetFollowers(userid int) []int {
	var follows []int
	sqlStr := `select followActfrom followTable 
	where userAct = ?`
	rows, err := db.Query(sqlStr, userid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var account int
		err := rows.Scan(&account)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
		follows = append(follows, account)
	}

	return follows
}

// 是否已经关注该用户
func IsFollowed(userid, account int) (bool, bool) {
	sqlStr := `select * from followTable 
	where userAct = ? and followAct=?`
	rows, err := db.Query(sqlStr, userid, account)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false, false
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		return true, true
	}
	return false, true
}

// 查找用户手机号
func FindPhone(id int) string { // 查找手机号是否存在
	var phoneNumber string
	sqlStr := `select phoneNumber
	from userinfo 
	where userAccount = ?`
	err := db.QueryRow(sqlStr, id).Scan(&phoneNumber)
	if err != nil {
		return phoneNumber
	} else {
		return phoneNumber
	}
}

func ModifyPassword(id int, secret string, password string) bool { //修改密码
	var secDB string
	sqlStr := `select password from userInfo where userAccount = ?`
	err := db.QueryRow(sqlStr, id).Scan(&secDB)
	if err != nil {
		return false
	}
	if secDB == secret {
		sqlStr = `update userInfo set password = ? where userAccount = ?`
		_, err = db.Exec(sqlStr, password, id)
		return err == nil
	} else {
		return false
	}
}
