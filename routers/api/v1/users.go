package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

type UsersInfo struct {
	Infos    models.UserInfo `json:"userInfo"` // 用户信息，只有一条，不用数组
	Notes    []models.Note   `json:"notes"`    // 笔记，简要信息
	Collects []models.Note   `json:"collects"`
	Likes    []models.Note   `json:"likes"`
	IsHost   bool            `json:"isHost"` //是否页面主人
}

type ModifiedInfo struct {
	Infos  models.ModifiableInfo `json:"userInfo"` //
	IsHost bool                  `json:"isHost"`   //是否页面主人
}

// 显示用户界面全部信息
func GetUserInfo(c *gin.Context) {
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

// 用户修改个人信息
func ModifyUserInfo(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	fmt.Println("用户ID:", userID)
	// 新声明的可修改信息的结构体
	var info ModifiedInfo
	info.Infos.Birthday = c.PostForm("birthday") // 从前端获取数据
	info.Infos.Gender = c.PostForm("gender")
	info.Infos.Introduction = c.PostForm("introduction")
	info.Infos.Password = c.PostForm("password")
	info.Infos.UserName = c.PostForm("userName")
	isHost := c.PostForm("isHost")
	info.IsHost, _ = strconv.ParseBool(isHost)

	files, err := GetHeadfile(userID) // 获取头像文件信息
	if err != nil {
		fmt.Println("用户没有以前的头像")
	}
	for _, file := range files {
		err := os.Remove(file) // 删除以前的头像
		if err != nil {
			return
		}
	}

	phone := models.FindPhone(userID) // 通过用户ID查找手机号

	file, _ := c.FormFile("file")
	log.Println(file.Filename)                                                               //输出文件名
	timeStamp := time.Now().Unix()                                                           // 时间戳
	name := fmt.Sprintf("head_%s_%s_%s", phone, strconv.Itoa(int(timeStamp)), file.Filename) // 文件名
	dst := fmt.Sprintf("images/%s", name)                                                    //路径
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst) // 图片
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	info.Infos.Portrait = name // 将图片目录保存在数据库

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

// 获取用户头像文件
func GetHeadfile(userid int) ([]string, error) {
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
	filter := fmt.Sprintf("head_%d", userid)
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), filter) {
			files = append(files, fmt.Sprintf("%s%s", "images/", file.Name()))
		}
	}
	return files, nil
}
