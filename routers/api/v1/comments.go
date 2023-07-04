package v1

import (
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

// 笔记页面加载评论
func GetComments(c *gin.Context) {
	success1 := true
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	comments, _, success := models.GetCommentInfo(noteId, 0)
	var temp bool
	//加载在评论的@信息
	for _, comment := range comments {
		comment.AtName, comment.AtLocation, temp = models.NewGetAtInfo(0, int(comment.CommentID))
		success1 = success1 && temp
	}
	if success && success1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    comments,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "fail",
			"data":    nil,
		})
	}
}

// 发表评论
func PostComment(c *gin.Context) {
	// var success bool
	// var success1 bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	//获取前端传来的数据
	var newComment models.Comment
	//通过ShouldBind获取json数据
	if err := c.ShouldBind(&newComment); err == nil {
		//获取@的信息
		newAt := make([]models.AtInfo, len(newComment.AtName), 50)
		for i := 0; i < len(newComment.AtName); i++ {
			newAt[i].AtName = newComment.AtName[i]
		}
		for j := 0; j < len(newComment.AtLocation); j++ {
			newAt[j].AtLocation = newComment.AtLocation[j]
		}
		cmtId, success := models.NewComment(newComment, noteId)
		success1 := models.AddAtInfo(int(newComment.CommentatorID), cmtId, 0, newAt)
		if success && success1 {
			//将该笔记点赞数加一
			models.ChangeNoteComments(noteId, 1)
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

// 删除评论
func CancleComment(c *gin.Context) {
	var success bool
	noteId, _ := strconv.Atoi(c.Param("noteId"))
	var comment models.Comment
	if err := c.ShouldBind(&comment); err == nil {
		success = models.DeleteComment(int(comment.CommentID))
		success1 := models.DeleteAtInfo(int(comment.CommentID))
		if success && success1 {
			//将该笔记点赞数加一
			models.ChangeNoteComments(noteId, -1)
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

// 消息列表加载评论
func MsgGetComments(c *gin.Context) {
	//是否全部已读，false表示没有
	totalState := true
	success1 := true
	userId, _ := strconv.Atoi(c.Param("userId"))
	comments, State, success := models.GetCommentInfo(userId, 1)
	//state表示整体的状态，0表示有未读的
	if State == 0 {
		totalState = false
	}
	var temp bool
	//加载在评论的@信息
	for _, comment := range comments {
		comment.AtName, comment.AtLocation, temp = models.NewGetAtInfo(0, int(comment.CommentID))
		success1 = success1 && temp
	}
	if success && success1 {
		c.JSON(http.StatusOK, gin.H{
			"code":       200,
			"message":    "success",
			"totalState": totalState,
			"data":       comments,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       400,
			"message":    "fail",
			"totalState": totalState,
			"data":       comments,
		})
	}
}

// 修改某条评论状态
func ChangeCommentState(c *gin.Context) {
	var success bool
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	success = models.SetCommentState(commentId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "评论状态修改失败",
		})
	}
}
