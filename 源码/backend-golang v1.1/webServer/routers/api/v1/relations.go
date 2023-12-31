package v1

//用来写评论、点赞、收藏相关的处理函数
//以上动作都是通过在关系数据库添加记录完成的
import (
	"fmt"
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

// 点赞某篇笔记
func LikeNote(c *gin.Context) {
	//数据库修改是否成功
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var likeInfo models.LikeInfo
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&likeInfo); err == nil {
		//向数据库中插入点赞信息
		success = models.NewLike(likeInfo, noteId)
		if success {
			//将该笔记点赞数加一
			models.ChangeNoteLikes(noteId, 1)
			//将该笔记作者点赞数加一
			models.ChangeUserLikes(noteId, 1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "点赞成功！",
			})
		} else {
			//取消前面的点赞信息插入（好像可以省下来？）
			models.DeleteLike(likeInfo, noteId)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "点赞失败！",
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

// 取消点赞
func CancelLike(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var likeInfo models.LikeInfo
	if err := c.ShouldBind(&likeInfo); err == nil {
		success = models.DeleteLike(likeInfo, noteId)
		if success {
			//将该笔记点赞数减一
			models.ChangeNoteLikes(noteId, -1)
			//将该笔记作者点赞数减一
			models.ChangeUserLikes(noteId, -1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "取消点赞成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "取消点赞失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}

// 收藏某篇笔记
func CollectNote(c *gin.Context) {
	//数据库修改是否成功
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var collectInfo models.CollectInfo
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&collectInfo); err == nil {
		//向数据库中插入收藏信息
		success = models.NewCollect(collectInfo, noteId)
		if success {
			//将该笔记收藏数加一
			models.ChangeNoteCollects(noteId, 1)
			//将该笔记作者收藏数加一
			models.ChangeUserCollects(noteId, 1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "收藏成功！",
			})
		} else {
			//取消前面的点赞信息插入（好像可以省下来？）
			models.DeleteCollect(collectInfo, noteId)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "收藏失败！",
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

// 取消收藏
func CancleCollect(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var collectInfo models.CollectInfo
	if err := c.ShouldBind(&collectInfo); err == nil {
		success = models.DeleteCollect(collectInfo, noteId)
		if success {
			//将该笔记收藏数减一
			models.ChangeNoteCollects(noteId, -1)
			//将该笔记作者收藏数减一
			models.ChangeUserCollects(noteId, -1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "取消收藏成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "取消收藏失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}

// 获取关注用户
func GetFollowUser(c *gin.Context) {
	var success bool
	var follows []models.FollowInfo
	userId, _ := strconv.Atoi(c.Param("userId"))
	follows, success = models.GetFollows(userId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    follows,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
		})
	}
}

// 关注用户
func FollowHandler(c *gin.Context) {
	var success bool
	userId, _ := strconv.Atoi(c.Param("userId"))
	var account models.FollowRequest
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&account); err == nil {
		id := account.FollowID
		fmt.Println(account)
		//向数据库中插入关注信息
		success = models.AddFollowInfo(userId, id)
		if success {
			//将用户关注数加一
			models.ChangeUserFollows(userId, 1)
			//将被关注用户粉丝数加一
			models.ChangeUserFans(id, 1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "关注成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "关注失败！",
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

// 取消关注
func CancelFollowHandler(c *gin.Context) {
	var success bool
	userId, _ := strconv.Atoi(c.Param("userId"))
	var account models.FollowRequest
	//用shouldBind获取前端传来的json数据，只要json名相同就能读取
	if err := c.ShouldBind(&account); err == nil {
		id := account.FollowID
		//向数据库中删除关注信息
		success = models.DelFollowInfo(userId, id)
		if success {
			//将用户关注数减一
			models.ChangeUserFollows(userId, -1)
			//将被关注用户粉丝数减一
			models.ChangeUserFans(id, -1)
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "取关成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "取关失败！",
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

// 消息列表加载点赞信息
func MsgGetLikes(c *gin.Context) {
	//是否全部已读，false表示没有
	totalState := true
	userId, _ := strconv.Atoi(c.Param("userId"))
	likes, State, success := models.GetLikeInfos(userId)
	if State == 0 {
		totalState = false
	}
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":       200,
			"message":    "success",
			"totalState": totalState,
			"data":       likes,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       400,
			"message":    "fail",
			"totalState": totalState,
			"data":       likes,
		})
	}
}

// 把点赞信息设为已读
func ChangeLikeState(c *gin.Context) {
	var success bool
	commentId, _ := strconv.Atoi(c.Param("fvId"))
	success = models.SetLikeState(commentId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "点赞状态修改失败",
		})
	}
}
