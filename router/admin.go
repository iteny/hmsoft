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
func AdminRouter(e *gin.Engine) {
	v1 := e.Group("/admin")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "login nima",
			})
		})
		v1.GET("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})
		v1.GET("/read", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})
	}
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
