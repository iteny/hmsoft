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
	// r.Use(Cors())
	router.AdminRouter(r)
	router.HomeRouter(r)

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

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method               //请求方法
// 		origin := c.Request.Header.Get("Origin") //请求头部
// 		var headerKeys []string                  // 声明请求头keys
// 		for k, _ := range c.Request.Header {
// 			headerKeys = append(headerKeys, k)
// 		}
// 		headerStr := strings.Join(headerKeys, ", ")
// 		if headerStr != "" {
// 			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
// 		} else {
// 			headerStr = "access-control-allow-origin, access-control-allow-headers"
// 		}
// 		originDomains := []string{"http://locahost:9090"}
// 		inArraysFlag := false
// 		for _, value := range originDomains {
// 			if origin == value {
// 				inArraysFlag = true
// 				break
// 			}
// 		}

// 		if origin != "" && inArraysFlag {
// 			c.Header("Access-Control-Allow-Origin", origin)                                    // 这是允许访问所有域
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
// 			//  header的类型
// 			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
// 			//              允许跨域设置                                                                                                      可以返回其他子段
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
// 			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
// 			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
// 			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
// 		}

// 		//放行所有OPTIONS方法
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "Options Request!")
// 		}
// 		// 处理请求
// 		c.Next() //  处理请求
// 	}
// }
