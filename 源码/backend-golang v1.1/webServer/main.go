// 主页面
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"webServer/models"

	"webServer/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 初始化数据库
	err := models.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	router := routers.InitRouter()

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8081),
		Handler:        router,
		MaxHeaderBytes: 8 << 20,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 在这里阻塞等待，当接收到上述两种信号时才会往下执行
	<-quit
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
