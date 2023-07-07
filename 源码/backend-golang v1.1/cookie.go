package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "1234", 3600, "/", "127.0.0.1", false, true) //
			/*
				第一个参数为 cookie 名；
				第二个参数为 cookie 值；
				第三个参数为 cookie 有效时长，当 cookie 存在的时间超过设定时间时，cookie 就会失效，它就不再是我们有效的 cookie；
				第四个参数为 cookie 所在的目录；
				第五个为所在域，表示我们的 cookie 作用范围；
				第六个表示是否只能通过 https 访问；
				第七个表示 cookie 是否可以通过 js代码进行操作。
			*/
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
