package v1

import (
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

// 加载评论
func GetComments(c *gin.Context) {
	var comments []models.Comment
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	comments, success = models.GetCommentInfo(noteId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    comments,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
			"data":    comments,
		})
	}
}

// 发表评论
func PostComment(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	//获取前端传来的数据
	var newComment models.Comment
	//通过ShouldBind获取json数据
	if err := c.ShouldBind(&newComment); err == nil {
		success = models.NewComment(newComment, noteId)
		if success {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "评论成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "评论失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}

func CancleComment(c *gin.Context) {
	var success bool
	// noteId, _ := strconv.Atoi(c.Param("noteId"))
	var comment models.Comment
	if err := c.ShouldBind(&comment); err == nil {
		success = models.DeleteComment(int(comment.CommentID))
		if success {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "评论删除成功！",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "取消评论失败！",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
}
