package webjwt

import (
	"github.com/gin-gonic/gin"
)

// IAuthorizator 授权规则接口
type IAuthorizator interface {
	HandleAuthorizator(data interface{}, c *gin.Context) bool
}

// UserAuthorizator 管理员授权规则
type UserAuthorizator struct {
}

// HandleAuthorizator 处理管理员授权规则
// func (*UserAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
// 	if v, ok := data.(*models.User); ok {
// 		for _, itemClaim := range v.UserClaims {
// 			if itemClaim.Type == "role" && itemClaim.Value == "admin" {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // TestAuthorizator 测试用户授权规则
// type TestAuthorizator struct {
// }

// // HandleAuthorizator 处理测试用户授权规则
// func (*TestAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
// 	if v, ok := data.(*models.User); ok && v.UserName == "test" {
// 		return true
// 	}
// 	return false
// }

// Visitor 游客授权规则
type Visitor struct {
}

// HandleAuthorizator 处理普通用户授权规则
func (*Visitor) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}
