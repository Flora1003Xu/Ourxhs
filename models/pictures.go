package models

import (
	"fmt"
)

type Pictures struct {
	PicID  string `json:"pic"`
	NoteId int    `json:"noteid"`
	Picurl string `json:"picurl"`
	Pictag string `json:"pictag"`
}

func NewPicInfo(pc Pictures) {
	sqlstr := `INSERT INTO pictureLibrary
	(noteID, picUrl, picTag)
	VALUES
	(?,?,?)`
	ret, err := db.Exec(sqlstr, pc.NoteId, pc.Picurl, pc.Pictag)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, affected rows:%d\n", theID)
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
