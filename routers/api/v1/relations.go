package v1

//用来写评论、点赞、收藏相关的处理函数
//以上动作都是通过在关系数据库添加记录完成的
import (
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
	// var success bool

}
