package v1

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"webServer/middleware/webjwt"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) { // 注册
	var requestUser models.Regist

	registTime := time.Now().Format("2006-01-02 15:04:05")

	requestUser.Birthday = c.PostForm("birthday") // 从前端获取数据
	requestUser.Gender = c.PostForm("gender")
	requestUser.Introduction = c.PostForm("introduction")
	requestUser.Password = c.PostForm("password")
	requestUser.PhoneNumber = c.PostForm("phoneNumber")
	requestUser.UserName = c.PostForm("userName")

	name := requestUser.UserName
	telephone := requestUser.PhoneNumber
	password := requestUser.Password

	//数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"data":    nil,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"data":    nil,
			"message": "密码不能少于6位",
		})
		return
	}
	if len(name) == 0 {
		var letters = []byte("abcdefghijklmnopqrstuvwxyz") // 随机生成用户名
		result := make([]byte, 10)
		for i := range result {
			result[i] = letters[rand.Intn(len(letters))]
		}
		name = string(result)
	}
	requestUser.UserName = name

	_, isExit := models.IsTelephoneExists(telephone)
	//判断手机号码是否存在
	if isExit { // 在数据库查找手机号码是否存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"data":    nil,
			"message": "用户已经存在",
		})
		return
	}

	file, _ := c.FormFile("file")
	log.Println(file.Filename)                                                                      //输出文件名
	timeStamp := time.Now().Unix()                                                                  // 时间戳
	PicName := fmt.Sprintf("head_%s_%s_%s", telephone, strconv.Itoa(int(timeStamp)), file.Filename) // 文件名
	dst := fmt.Sprintf("images/%s", PicName)                                                        //路径
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst) // 图片
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	requestUser.Portrait = PicName // 将图片目录保存在数据库

	//把上述的数据存入数据库，从而创建新用户
	if !models.CreateUser(requestUser, registTime) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"data":    nil,
			"message": "数据库写入失败",
		})
	} else {
		//返回结果
		//发放token
		token, err := webjwt.ReleaseToken(requestUser.PhoneNumber)
		if err != nil { // token发放失败
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    500,
				"data":    nil,
				"message": "系统异常",
			})
			log.Printf("token generate error: %v", err)
			return
		}

		//返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    token, // data中存放token
			"message": "注册成功",
		})
	}
}

func Login(c *gin.Context) {
	var requestUser = models.Regist{}
	c.Bind(&requestUser)

	//获取参数
	telephone := requestUser.PhoneNumber
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"data":    nil,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"data":    nil,
			"message": "密码不能少于6位",
		})
		return
	}

	userID, isExit := models.IsTelephoneExists(telephone) // 是否存在

	//判断手机号码是否存在
	if !isExit { // 在数据库查找手机号码是否存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    400,
			"data":    nil,
			"message": "用户不存在",
		})
		return
	}

	if !models.SecretCorrect(telephone, password) { // 在数据库查找手机号码是否存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    400,
			"data":    nil,
			"message": "密码错误", // 前端提示信息
		})
		return
	}

	//发放token
	token, err := webjwt.ReleaseToken(requestUser.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"data":    nil,
			"message": "系统异常",
		})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    token,
		"uid":     userID,
		"message": "登录成功",
	})
}

// // 从上下文中获取用户信息
// func Info(ctx *gin.Context) {
// 	user, _ := ctx.Get("user")
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code": 200,
// 		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
// 	})
// }
