package models

//点赞、收藏、关注中与数据库相关的操作

import (
	"fmt"
)

// 点赞信息
type LikeInfo struct {
	FvID        int    `json:"fvId" form:"fvId"`
	UserAct     int    `json:"userAct" form:"userAct"`
	FavorNoteID int    `json:"favorNoteId" form:"favorNoteId"`
	State       int    `json:"state" form:"state"`
	LikeTime    string `json:"likeTime" form:"likeTime"`
}

// 消息列表的笔记信息
type LikeToShow struct {
	LikeID   int    `json:"likeId" form:"likeId"`
	UserName string `json:"userName" form:"userName"`
	Portrait string `json:"portrait" form:"portrait"`
	LikeTime string `json:"likeTime" form:"likeTime"`
	NoteID   int    `json:"noteId" form:"noteId"`
	State    int    `json:"state" form:"state"`
}

// 收藏信息
type CollectInfo struct {
	CltID         int `json:"cltId" form:"cltId"`
	UserAct       int `json:"userAct" form:"userAct"`
	CollectNoteID int `json:"collectNoteId" form:"collectNoteId"`
}

// @用户信息
type AtInfo struct {
	AtUserID   int    `json:"atUserId" form:"atUserId"`
	AtName     string `json:"atName" form:"atName"`
	AtLocation int    `json:"atLocation" form:"atLocation"`
	AtState    int    `json:"atState" form:"atState"`
}

// 要显示的关注的人信息
type FollowInfo struct {
	FolInfoId int    `json:"folInfoId" form:"folInfoId"`
	FollowAct int    `json:"followAct" form:"followAct"`
	UserName  string `json:"userName" form:"userName"`
	Portrait  string `json:"portrait" form:"portrait"`
}

// 关注信息
type FollowRequest struct {
	FollowID int `json:"followID" form:"followID"` // 关注的人的id
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
	sqlstr := `INSERT INTO favorTable (userAct, favorNoteId, fvTime) VALUES (?,?,?)`
	ret, err := db.Exec(sqlstr, nl.UserAct, noteId, nl.LikeTime)
	if err != nil {
		fmt.Printf("点赞信息insert failed, err:%v\n", err)
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
func GetFollows(userId int) (follows []FollowInfo, ok bool) {
	ok = true
	sqlstr := `SELECT f.folInfoId, f.followAct, u.userName, u.portrait
	FROM followTable f, userInfo u
	WHERE f.userAct=? AND f.followAct=u.userAccount `
	rows, err := db.Query(sqlstr, userId)
	if err != nil {
		fmt.Printf("获取关注用户query failed, err:%v\n", err)
		ok = false
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var fl FollowInfo
		err := rows.Scan(&fl.FolInfoId, &fl.FollowAct, &fl.UserName, &fl.Portrait)
		if err != nil {
			fmt.Printf("获取关注用户scan failed, err:%v\n", err)
			ok = false
			return
		}
		follows = append(follows, fl)
	}
	return follows, ok
}

// 写入某个东西的@信息
func AddAtInfo(userId, noteId, state int, atinfos []AtInfo) bool {
	//state字段表示是在正文还是评论，1表示在正文，0表示在评论
	sqlstr := "INSERT INTO atTable (userAct, noteId, atUserName, atLocation, atState) VALUES (?,?,?,?,?)"
	for _, atItem := range atinfos {
		_, err := db.Exec(sqlstr, atItem.AtUserID, noteId, atItem.AtName, atItem.AtLocation, state)
		if err != nil {
			fmt.Printf("@信息insert failed, err:%v\n", err)
			return false
		}
	}
	return true
}

// 删除某个东西（笔记或评论）的@信息
func DeleteAtInfo(noteId int) bool {
	sqlstr := "DELETE from atTable WHERE noteId=?"
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("@信息delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("@信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("@信息 delete success, affected rows:%d\n", n)
	return true
}

// 消息列表获取点赞信息
func GetLikeInfos(userId int) (likeInfos []LikeToShow, totalState int, ok bool) {
	ok = true
	sqlstr := `SELECT f.fvId, f.state, f.fvTime, u.userName, u.portrait
	FROM favorTable f, userInfo u, noteInfo n
	WHERE n.creatorAccount=? AND n.noteId=f.favorNoteId AND f.userAct=u.userAccount
	ORDER BY f.fvTime DESC`
	rows, err := db.Query(sqlstr, userId)
	if err != nil {
		fmt.Printf("点赞query failed, err:%v\n", err)
		ok = false
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	//整体的状态，1表示已读
	totalState = 1
	// 循环读取结果集中的数据
	for rows.Next() {
		var lk LikeToShow
		err := rows.Scan(&lk.LikeID, &lk.State, &lk.LikeTime, &lk.UserName, &lk.Portrait)
		if err != nil {
			fmt.Printf("点赞scan failed, err:%v\n", err)
			ok = false
			return
		}
		lk.LikeTime = JudgeTime(lk.LikeTime)
		totalState = totalState * lk.State
		likeInfos = append(likeInfos, lk)
	}
	return likeInfos, totalState, ok
}

// 增加关注信息
func AddFollowInfo(useraccount int, account int) bool {
	sqlstr := "insert into followTable (userAct,followAct) values (?,?)"
	ret, err := db.Exec(sqlstr, useraccount, account)
	if err != nil {
		fmt.Printf("关注信息insert failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("关注信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("关注信息插入 success, affected rows:%d\n", n)
	return true
}

// 删除关注信息
func DelFollowInfo(useraccount int, account int) bool {
	sqlstr := "DELETE from followTable where userAct=? and followAct=?"
	ret, err := db.Exec(sqlstr, useraccount, account)
	if err != nil {
		fmt.Printf("关注信息delete failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("关注信息get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("关注信息 delete success, affected rows:%d\n", n)
	return true
}

// 修改关注数
func ChangeUserFollows(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set followNum = followNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set followNum =followNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("关注数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("关注数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("关注数修改编号：%d\n", n)
}

// 修改粉丝数
func ChangeUserFans(userId, option int) {
	var sqlstr string
	addnum := `UPDATE userInfo set fansNum =fansNum+1 WHERE userAccount = ?`
	reducenum := `UPDATE userInfo set fansNum =fansNum-1 WHERE userAccount = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, userId)
	if err != nil {
		fmt.Printf("粉丝数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("粉丝数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("粉丝数修改编号：%d\n", n)
}

// 加载@的信息
func NewGetAtInfo(state, id int) (atname []string, atid, atlocation []int, ok bool) {
	ok = true
	sqlstr := `SELECT userAct, atUserName, atLocation FROM atTable WHERE noteId=? and atState=?`
	rows, err := db.Query(sqlstr, id, state)
	if err != nil {
		ok = false
		fmt.Printf("@信息query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var at AtInfo
		err := rows.Scan(&at.AtUserID, &at.AtName, &at.AtLocation)
		if err != nil {
			ok = false
			fmt.Printf("@信息scan failed, err:%v\n", err)
			return
		}
		atid = append(atid, at.AtUserID)
		atname = append(atname, at.AtName)
		atlocation = append(atlocation, at.AtLocation)
	}
	return
}

// 把某个点赞设为已读
func SetLikeState(likeId int) bool {
	sqlstr := "UPDATE favorTable SET state=1 WHERE fvId=?"
	ret, err := db.Exec(sqlstr, likeId)
	if err != nil {
		fmt.Printf("点赞信息状态update failed, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("点赞状态get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("点赞状态修改编号：%d\n", n)
	return true
}

// 删除点赞信息
func RemoveLikes(noteId int) bool {
	sqlstr := "DELETE FROM favorTable WHERE favorNoteId=?"
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("点赞批量删除失败, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("点赞批量删除Get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("点赞信息批量 delete success, affected rows:%d\n", n)
	return true
}

// 批量删除收藏信息
func RemoveCollect(noteId int) bool {
	sqlstr := "DELETE FROM collectTable WHERE collectNoteId=?"
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("收藏批量删除失败, err:%v\n", err)
		return false
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("收藏批量删除Get RowsAffected failed, err:%v\n", err)
		return false
	}
	fmt.Printf("收藏信息批量 delete success, affected rows:%d\n", n)
	return true
}
