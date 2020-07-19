package main

import (
	"context"
	"hmsoft/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	time.Sleep(5 * time.Second)
	// 	c.String(http.StatusOK, "Welcome Gin Server")
	// })
	// r := router.AdminRouter()
	// r =router.HomeRouter()
	r := gin.Default()
	router.LoadBlog(r)
	router.LoadShop(r)
	srv := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
