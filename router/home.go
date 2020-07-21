package router

import "github.com/gin-gonic/gin"

//后台管理路由
// func HomeRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/home", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "home",
// 		})
// 	})
// 	// 更多的注册
// 	// r.POST()
// 	// r.DELETE()
// 	// r.PUT()
// 	return r
// }
func HomeRouter(e *gin.Engine) {
	e.GET("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})
  e.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})
}
