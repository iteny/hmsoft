package router

import "github.com/gin-gonic/gin"

//后台管理路由
// func AdminRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	// 更多的注册
// 	// r.POST()
// 	// r.DELETE()
// 	// r.PUT()
// 	return r
// }
func LoadShop(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	e.GET("/good", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "good",
		})
	})
	e.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
}
