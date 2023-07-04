package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type Data struct {
	IsLogin bool          `json:"isLogin"` // 是否登录
	Notes   []models.Note `json:"notes"`   // 笔记，简要信息
}
type Note models.Note

// 获取笔记（全部）
func GetAllNotes(c *gin.Context) {
	var data Data
	var ok bool
	//判断是否登录，还要再加判断的函数
	data.IsLogin = false
	data.Notes, ok = models.GetBriefNtInfo()
	// gin.H 是map[string]interface{}的缩写
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    data,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
			"data":    data,
		})
	}
}

// 获取特定笔记（搜索/标签）
func GetSpecificNotes(c *gin.Context) {
	var data Data
	var OK bool
	//判断是否登录，还要再加判断的函数
	data.IsLogin = false
	keyword := c.Param("keyword")
	data.Notes, OK = models.GetSpBriefNtInfo(keyword)
	if OK {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    data,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
			"data":    data,
		})
	}
}

// 获取笔记详细内容
func NoteDetailHandler(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userId"))
	noteid, _ := strconv.Atoi(c.Param("noteid"))
	data := models.SpecificNote(noteid)
	data.NoteInfo.IsCollected = models.IsCollected(userid, noteid)
	authorid := models.NoteToUser(noteid)
	data.NoteInfo.IsFollowed = models.IsFollowed(userid, authorid)
	data.NoteInfo.IsLiked = models.IsLiked(userid, noteid)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// 获取关注的人的笔记
func GetFollowedNotes(c *gin.Context) {
	var data Data
	var OK bool
	//判断是否登录，还要再加判断的函数
	data.IsLogin = false
	userId, _ := strconv.Atoi(c.Param("userId"))
	data.Notes, OK = models.GetFlwedNotes(userId)
	if OK {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    data,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "获取关注人笔记失败！",
			"data":    data,
		})
	}
}

// 上传笔记
func UploadNote(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": fmt.Sprintf("读取失败! err:%s", err),
		})
		return
	} else {
		files := form.File["files"]

		//新声明新笔记的结构体
		var newNote models.DetailNote
		fmt.Print("最初的笔记信息")
		fmt.Println(newNote)
		if err := c.ShouldBind(&newNote); err == nil {
			newNote.CreatorID = userId
			//先创建一个不含内容的笔记信息，后面再填上信息
			fmt.Print("初始的笔记信息")
			fmt.Println(newNote)
			fmt.Printf("Tag里内容:%d", len(newNote.Tags))
			fmt.Println(newNote.Tags)
			ntID, success := models.NewNoteInfo(newNote)
			if !success {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": "笔记上传失败（数据库创建失败）",
				})
				return
			}

			fileType := map[string]bool{
				".png":  true,
				".jpg":  true,
				".jpeg": true,
				".gif":  true,
			}
			for index, file := range files {
				fmt.Print(index)
				extName := path.Ext(file.Filename)
				_, b := fileType[extName]
				if !b {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    400,
						"message": "上传文件类型不合法！",
					})
					return
				}
				var pc models.Pictures
				timeStamp := time.Now().Unix()

				log.Println(file.Filename)
				name := fmt.Sprintf("%d_%d_%s_%s", userId, ntID, strconv.Itoa(int(timeStamp)), file.Filename)
				dst := fmt.Sprintf("images/%s", name)
				if index == 0 {
					newNote.Cover = name
				}
				pc.NoteId = ntID
				pc.Picurl = name
				// 上传文件到指定的目录
				c.SaveUploadedFile(file, dst)
				//将路径等信息更新到数据库
				models.NewPicInfo(pc)
			}
			newNote.NoteID = ntID
			newNote.Picnum = len(files)
			newAt := make([]models.AtInfo, len(newNote.AtName), 50)
			for i := 0; i < len(newNote.AtName); i++ {
				newAt[i].AtName = newNote.AtName[i]
			}
			for j := 0; j < len(newNote.AtLocation); j++ {
				newAt[j].AtLocation = newNote.AtLocation[j]
			}
			//写入笔记的全部信息
			fmt.Print("完整的笔记信息")
			fmt.Println(newNote)
			ok := models.ModifyNote(newNote)
			//写入@信息
			ok1 := models.AddAtInfo(userId, ntID, 1, newAt)
			if ok && ok1 {
				models.ChangeNoteNum(userId, 1)
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": fmt.Sprintf("笔记上传成功，共%d张图片", len(files)),
				})
			} else {
				//如果上传失败，就把空的信息删掉
				models.DeleteNoteInfo(userId)
				models.DeletePic(ntID)
				models.DeleteAtInfo(ntID)
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": "笔记上传失败（数据库写入失败）",
				})
			}
		} else {
			//json数据获取失败
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  400,
				"error": err.Error(),
			})
		}
	}
}

// 删除笔记
func DeleteNote(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	files, err := Getfile(userId, noteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "要删除的笔记/图片不存在！",
		})
		return
	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "删除失败",
			})
			return
		}
	}
	if models.DeletePic(noteId) && models.DeleteNoteInfo(noteId) && models.RemoveComments(noteId) && models.RemoveLikes(noteId) {
		models.ChangeNoteNum(userId, 0)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除成功！",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "数据库删除失败！",
		})
	}
}

// 获取笔记对应的图片文件
func Getfile(userid, noteid int) ([]string, error) {
	var files []string
	f, err := os.Open("images")
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	filter := fmt.Sprintf("%d_%d", userid, noteid)
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), filter) {
			files = append(files, fmt.Sprintf("%s%s", "images/", file.Name()))
		}
	}
	return files, nil
}

// // 获取关注的人的笔记
// func GetFollowersNotesHandler(ctx gin.Context) {
// 	// 解析用户id
// 	userid, _ := strconv.Atoi(ctx.Param("userId"))
// 	// 获取该用户的关注人列表
// 	var follows []int = models.GetFollowers(userid)
// 	// 根据(批量)用户id查找笔记

// }

// 获取走马灯的4个笔记
func GetTops(c *gin.Context) {
	notes, success := models.Tops()
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    notes,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
			"data":    notes,
		})
	}
}
