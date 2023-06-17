package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type UsersInfo struct {
	Infos    models.UserInfo `json:"userInfo"` // 用户信息，只有一条，不用数组
	Notes    []models.Notes  `json:"notes"`    // 笔记，简要信息
	Collects []models.Notes  `json:"collects"`
	Likes    []models.Notes  `json:"likes"`
	IsHost   bool            `json:"isHost"` //是否页面主人
}

type ModifiedInfo struct {
	Infos  models.ModifiableInfo `json:"userInfo"` //
	IsHost bool                  `json:"isHost"`   //是否页面主人
}

func GetUserInfo(c *gin.Context) { //显示用户界面全部信息
	var info UsersInfo
	userID, _ := strconv.Atoi(c.Param("userId"))
	fmt.Println("用户ID:", userID)
	// 通过用户ID去数据库获取信息
	info.Infos = models.UserInfoDB(userID)
	// 获取某用户发布的笔记
	info.Notes = models.NoteInfoDB(userID)
	// 获取某用户收藏的笔记
	info.Collects = models.CollectInfoDB(userID)
	// 获取某用户点赞的笔记
	info.Likes = models.LikeInfoDB(userID)
	info.IsHost = true
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    info,
	})
}

func ModifyUserInfo(c *gin.Context) { //用户修改个人信息
	userID, _ := strconv.Atoi(c.Param("userId"))
	fmt.Println("用户ID:", userID)
	// 新声明的可修改信息的结构体
	var info ModifiedInfo
	info.Infos.Birthday = c.PostForm("birthday") // 从前端获取数据
	info.Infos.Gender = c.PostForm("gender")
	info.Infos.Introduction = c.PostForm("introduction")
	info.Infos.Mail = c.PostForm("mail")
	info.Infos.Password = c.PostForm("password")
	info.Infos.PhoneNumber = c.PostForm("phoneNumber")
	info.Infos.Portrait = c.PostForm("portrait")
	info.Infos.UserName = c.PostForm("userName")
	info.IsHost, _ = strconv.ParseBool(c.PostForm("isHost"))
	success := models.ModifyInfo(info.Infos, userID)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "表单内容获取失败!",
			"data":    info, // 是否需要重新返回呢，不需要则去掉data字段
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    info, // 将修改的数据发送回前端
	})
}
