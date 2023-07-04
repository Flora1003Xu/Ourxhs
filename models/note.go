package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Note struct {
	Cover       string `json:"cover"`
	CreatorID   int    `json:"creatorId"`
	CreatorName string `json:"creatorName"` // 作者编号
	LikedNum    int    `json:"likedNum"`    // 点赞数
	NoteID      int    `json:"noteId"`      // 笔记编号
	Portrait    string `json:"portrait"`    // 头像
	Title       string `json:"title"`
}

// 笔记的详细信息
type DetailNote struct {
	NoteID     int        `json:"noteid" form:"noteid"`       // 笔记编号
	CreatorID  int        `json:"creatorId" form:"creatorId"` // 作者编号
	Title      string     `json:"title" form:"title"`
	Body       string     `json:"body" form:"body"`
	Picnum     int        `json:"picnum" form:"picnum"`
	Cover      string     `json:"cover" form:"cover"`
	CreateTime string     `json:"createtime" form:"createtime"`
	UpdateTime string     `json:"updatetime" form:"updatetime"`
	Tags       [11]string `json:"tags" form:"tags"`
	Location   string     `json:"location" form:"location"`
	AtUserID   int        `json:"atuserid" form:"atuserid"`
	LikedNum   int        `json:"likedNum" form:"likedNum"` // 点赞数  待定
	AtName     []string   `json:"atName" form:"atName"`
	AtLocation []int      `json:"atLocation" form:"atLocation"`

	IsCollected bool `json:"isCollected"` // 是否收藏该篇
	IsFollowed  bool `json:"isFollowed"`  // 是否关注该作者
	IsLiked     bool `json:"isLiked"`     // 是否点赞该篇
	CollectNum  int  `json:"collectNum"`  // 收藏数
}

type fullNote struct {
	NoteInfo   DetailNote `JSON:"noteInfo"`
	PicsOfNote []string   `JSON:"pictures"`
}

// 获取笔记的封面标题等简要信息
func GetBriefNtInfo() (notes []Note, ok bool) {
	ok = true
	sqlStr := `select n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName
	from noteInfo n,userInfo u
	where n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr)
	if err != nil {
		ok = false
		fmt.Printf("query failed, err:%v\n", err)
		return notes, ok
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait, &nt.CreatorName)
		if err != nil {
			ok = false
			fmt.Printf("scan failed, err:%v\n", err)
			return notes, ok
		}
		fmt.Print(nt.NoteID)
		notes = append(notes, nt)
	}
	return notes, ok
}

// 获取特定内容的笔记
func GetSpBriefNtInfo(keyword string) (notes []Note, ok bool) {
	sqlStr := `SELECT n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName
	FROM noteInfo n, userInfo u
	WHERE n.creatorAccount = u.userAccount AND (n.tag = ? OR n.title LIKE CONCAT('%',#{keyword},'%'))`
	rows, err := db.Query(sqlStr, keyword, keyword)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil, false
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait, &nt.CreatorName)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, false
		}
		fmt.Println(nt.NoteID)
		notes = append(notes, nt)
	}
	return notes, true
}

// 获取用户关注的人的所有笔记的简要信息
func GetFlwedNotes(userId int) (notes []Note, ok bool) {
	sqlStr := `SELECT n.noteId, n.title, n.cover, n.creatorAccount, n.likeNum, u.portrait, u.userName
	FROM noteInfo n, userInfo u ,followTable f
	WHERE f.userAct=? AND f.followAct=n.creatorAccount AND n.creatorAccount = u.userAccount`
	rows, err := db.Query(sqlStr, userId)
	if err != nil {
		fmt.Printf("关注的人的笔记query failed, err:%v\n", err)
		return nil, false
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover, &nt.CreatorID, &nt.LikedNum, &nt.Portrait, &nt.CreatorName)
		if err != nil {
			fmt.Printf("关注的人的笔记scan failed, err:%v\n", err)
			return nil, false
		}
		fmt.Println(nt.NoteID)
		notes = append(notes, nt)
	}
	return notes, true
}

// 存入新上传的笔记信息
func NewNoteInfo(nn DetailNote) (int, bool) {
	sqlstr := `INSERT INTO noteInfo
	(creatorAccount, cover, title, body, createTime, updateTime, location, atUserId)
	VALUES
	(?,?,?,?,?,?,?,?)`
	ret, err := db.Exec(sqlstr, nn.CreatorID, nn.Cover, nn.Title, nn.Body, nn.CreateTime, nn.UpdateTime, nn.Location, nn.AtUserID)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return -1, false
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return -1, false
	}
	return int(theID), true
}

// 更新某笔记的信息
func ModifyNote(mn DetailNote) bool {
	sqlstr := `UPDATE noteInfo SET
	cover=?, title=?, body=?, numOfPic=?, createTime=?, updateTime=?, location=?, tag=?, tag1=?, tag2=?, tag3=?, tag4=?, tag5=?, tag6=?, tag7=?, tag8=?, tag9=?, tag10=?
	WHERE
	noteId=?`
	ret, err := db.Exec(sqlstr, mn.Cover, mn.Title, mn.Body, mn.Picnum, mn.CreateTime, mn.UpdateTime, mn.Location,
		mn.Tags[0], mn.Tags[1], mn.Tags[2], mn.Tags[3], mn.Tags[4], mn.Tags[5], mn.Tags[6], mn.Tags[7], mn.Tags[8], mn.Tags[9], mn.Tags[10],
		mn.NoteID)
	if err != nil {
		fmt.Printf("笔记修改 failed, err:%v\n", err)
		return false
	}
	theID, err := ret.LastInsertId() // 修改的行数
	if err != nil {
		fmt.Printf("get lastinsert ID failed,  err:%v\n", err)
		return false
	}
	fmt.Printf("笔记更新成功！，影响行数：%d\n", theID)
	return true
}

// 删除笔记
func DeleteNoteInfo(ntid int) bool {
	sqlstr := "DELETE FROM noteInfo WHERE noteID = ?"
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

// 返回笔记详情页
func SpecificNote(noteid int) fullNote {
	var N fullNote
	//先找笔记信息
	sqlStr1 := "select Noteid,CreatorAccount,Title,Body,NumOfPic,Cover,CreateTime,UpdateTime,Location,AtUserid from noteInfo where noteId = ?"
	rows, err := db.Query(sqlStr1, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		var err fullNote
		return err
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var createTimestring string
		var updateTimestring string
		err := rows.Scan(&N.NoteInfo.NoteID, &N.NoteInfo.CreatorID, &N.NoteInfo.Title, &N.NoteInfo.Body, &N.NoteInfo.Picnum, &N.NoteInfo.Cover, &createTimestring, &updateTimestring, &N.NoteInfo.Location, &N.NoteInfo.AtUserID)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			var err fullNote
			return err
		}
		N.NoteInfo.CreateTime = createTimestring
		N.NoteInfo.UpdateTime = updateTimestring
		// N.NoteInfo.CreateTime, _ = time.ParseInLocation("2000-01-01 24:00:00", createTimestring, time.Local)
		// N.NoteInfo.UpdateTime, _ = time.ParseInLocation("2000-01-01 24:00:00", updateTimestring, time.Local)
	}
	//再找图片信息
	sqlStr2 := "select picUrl from pictureLibrary where noteID = ?"
	rows2, err := db.Query(sqlStr2, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		var err fullNote
		return err
	}
	for rows2.Next() {
		var picurl string
		err := rows2.Scan(&picurl)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			var err fullNote
			return err
		}
		N.PicsOfNote = append(N.PicsOfNote, picurl)
	}

	// 关闭rows释放持有的s数据库链接
	defer rows2.Close()
	N.NoteInfo.CollectNum = GetNoteCollectNum(noteid)
	return N
}

// 修改笔记的获赞数
func ChangeNoteLikes(noteId, option int) {
	var sqlstr string
	addnum := `UPDATE noteInfo set likeNum =likeNum+1 WHERE noteId = ?`
	reducenum := `UPDATE noteInfo set likeNum =likeNum-1 WHERE noteId = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, noteId)
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
	fmt.Printf("笔记编号：%d\n", n)
}

// 修改笔记被收藏数
func ChangeNoteCollects(noteId, option int) {
	var sqlstr string
	addnum := `UPDATE noteInfo set collectNum =collectNum+1 WHERE noteId = ?`
	reducenum := `UPDATE noteInfo set collectNum =collectNum-1 WHERE noteId = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("笔记收藏数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("笔记收藏数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("笔记收藏数修改编号：%d\n", n)
}

// 修改笔记评论数
func ChangeNoteComments(noteId, option int) {
	var sqlstr string
	addnum := `UPDATE noteInfo set commentNum =commentNum+1 WHERE noteId = ?`
	reducenum := `UPDATE noteInfo set commentNum =commentNum-1 WHERE noteId = ?`
	if option == 1 {
		sqlstr = addnum
	} else {
		sqlstr = reducenum
	}
	ret, err := db.Exec(sqlstr, noteId)
	if err != nil {
		fmt.Printf("笔记评论数update failed, err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("笔记评论数get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("笔记评论数修改编号：%d\n", n)
}

// 获取笔记收藏数
func GetNoteCollectNum(noteid int) int {
	var collectNum int = 0
	sqlStr := "select * from collectTable where collectNoteId = ?"
	rows, err := db.Query(sqlStr, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return 0
	}
	for rows.Next() {
		collectNum += 1
	}
	defer rows.Close()
	return collectNum
}

// 判断是否收藏该笔记
func IsCollected(userid, noteid int) bool {
	sqlStr := "select * from collectTable where userAct=? and collectNoteId = ?"
	rows, err := db.Query(sqlStr, userid, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	for rows.Next() {
		return true
	}
	defer rows.Close()
	return false
}

// 判断是否收藏该笔记
func IsLiked(userid, noteid int) bool {
	sqlStr := "select * from favorTable where userAct=? and favorNoteId = ?"
	rows, err := db.Query(sqlStr, userid, noteid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	for rows.Next() {
		return true
	}
	defer rows.Close()
	return false
}

// 获取走马灯的笔记
func Tops() (notes []Note, ok bool) {
	ok = true
	sqlStr := `SELECT noteId, title, cover
	FROM noteInfo
	ORDER By likeNum LIMIT 4`
	rows, err := db.Query(sqlStr)
	if err != nil {
		ok = false
		fmt.Printf("走马灯query failed, err:%v\n", err)
		return notes, ok
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var nt Note
		err := rows.Scan(&nt.NoteID, &nt.Title, &nt.Cover)
		if err != nil {
			ok = false
			fmt.Printf("走马灯scan failed, err:%v\n", err)
			return notes, ok
		}
		fmt.Print(nt.NoteID)
		notes = append(notes, nt)
	}
	return notes, ok
}
