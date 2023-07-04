// 主页面
package main

import (
	"fmt"

	//"github.com/gin-gonic/gin"

	"webServer/models"

	"webServer/routers"

	_ "github.com/go-sql-driver/mysql"
)

// func hostFunc(c *gin.Context) {
// 	// gin.H 是map[string]interface{}的缩写
// 	c.HTML(http.StatusOK, "host.html", gin.H{
// 		"title": "标题",
// 		"nt":    models.Notes,
// 	})
// }
//   已转到routers/api/v1/notes.go

func main() {
	// 初始化数据库
	err := models.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	router := routers.InitRouter()

	router.Run(":8080")

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", 8080),
	// 	Handler:        router,
	// 	ReadTimeout:    60,
	// 	WriteTimeout:   60,
	// 	MaxHeaderBytes: 8 << 20,
	// }
	// s.ListenAndServe()

	// //加载HTML文件
	// router.LoadHTMLGlob("templates/host.html")
	// router.GET("/", hostFunc)
	// router.Run(":8080")
	// 已转到routers/router.go
}
