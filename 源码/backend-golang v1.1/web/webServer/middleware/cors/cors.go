package cors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 用于连接的中间件
// func CorsMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
// 		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		// //放行所有OPTIONS方法
// 		// if method == "OPTIONS" {
// 		// 	c.JSON(http.StatusOK, "Options Request!")
// 		// }
// 		// // 处理请求
// 		// c.Next() //  处理请求
// 		if ctx.Request.Method == http.MethodOptions {
// 			ctx.AbortWithStatus(200)
// 		} else {
// 			ctx.Next()
// 		}
// 	}
// }

// 用于跨区域连接的中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求方法
		method := c.Request.Method
		//请求头部
		origin := c.Request.Header.Get("Origin")
		// 声明请求头keys
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// 接收客户端发送的origin
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 这是允许访问所有域
			c.Header("Access-Control-Allow-Origin", "*")
			// 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//允许跨域设置可以返回其他子段，可以自定义字段	 // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			// 缓存请求信息 单位为秒
			c.Header("Access-Control-Max-Age", "172800")
			// 跨域请求是否需要带cookie信息 默认设置为true
			c.Header("Access-Control-Allow-Credentials", "false")
			// 设置返回格式是json
			c.Set("content-type", "application/json")
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
