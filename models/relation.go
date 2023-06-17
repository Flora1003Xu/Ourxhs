package models

import (
	"fmt"
)

// 点赞信息
type LikeInfo struct {
	FvID        int `json:"fvId"`
	UserAct     int `json:"userAct"`
	FavorNoteID int `json:"favorNoteId"`
}

// 收藏信息
type CollectInfo struct {
	CltID         int `json:"cltId"`
	UserAct       int `json:"userAct"`
	CollectNoteID int `json:"collectNoteId"`
}

// 根据笔记编号获取作者账号
func NoteToUser(noteId int) int {
	var userId int
	sltsql := "SELECT creatorAccount FROM noteInfo WHERE noteId=?"
	err := db.QueryRow(sltsql, noteId).Scan(&userId)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return -1
	} else {
		return userId
	}
}

// 插入点赞信息
func NewLike(nl LikeInfo, noteId int) bool {
	sqlstr := `INSERT INTO favorTable (userAct, favorNoteId) VALUES (?,?)`
	ret, err := db.Exec(sqlstr, nl.UserAct, noteId)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("点赞成功！评论在数据库行数：%d\n", theID)
	return true
}

// 删除点赞信息
func DeleteLike(nl LikeInfo, noteId int) bool {
	// userId := NoteToUser(noteId)
	sqlstr := "DELETE from favorTable WHERE userAct=? AND favorNoteId=?"
	ret, err := db.Exec(sqlstr, nl.UserAct, noteId)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("点赞信息 delete success, affected rows:%d\n", n)
	return true
}

// 插入收藏信息
func NewCollect(nclt CollectInfo, noteId int) bool {
	sqlstr := `INSERT INTO collectTable (userAct, collectNoteId) VALUES (?,?)`
	ret, err := db.Exec(sqlstr, nclt.UserAct, noteId)
	if err != nil {
		fmt.Printf("收藏信息insert failed, err:%v\n", err)
		return false
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("收藏信息get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏成功！收藏在数据库行数：%d\n", theID)
	return true
}

// 删除收藏信息
func DeleteCollect(dclt CollectInfo, noteId int) bool {
	sqlstr := "DELETE from collectTable WHERE userAct=? AND collectNoteId=?"
	ret, err := db.Exec(sqlstr, dclt.UserAct, noteId)
	if err != nil {
		fmt.Printf("收藏信息delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("收藏信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏信息 delete success, affected rows:%d\n", n)
	return true
}

// 获取关注用户
func GetFollows() bool {
	return true
}
