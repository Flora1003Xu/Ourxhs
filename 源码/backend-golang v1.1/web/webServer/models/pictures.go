package models

import (
	"fmt"
)

type Pictures struct {
	PicID  string `json:"pic"`
	NoteId int    `json:"noteid"`
	Picurl string `json:"picurl"`
}

func NewPicInfo(pc Pictures) bool {
	sqlstr := `INSERT INTO pictureLibrary
	(noteID, picUrl)
	VALUES
	(?,?)`
	ret, err := db.Exec(sqlstr, pc.NoteId, pc.Picurl)
	if err != nil {
		fmt.Printf("笔记图片insert failed, err:%v\n", err)
		return false
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("笔记图片insert success, affected rows:%d\n", theID)
	return true
}

func DeletePic(ntid int) bool {
	sqlstr := "DELETE FROM pictureLibrary WHERE noteID = ?"
	ret, err := db.Exec(sqlstr, ntid)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
	return true
}
