package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
		v1.POST("/login", func(c *gin.Context) {
			username := c.PostForm("username")
			// password := c.DefaultPostForm("nick", "password") // 此方法可以设置默认值
			fmt.Println(username + "1111")
			c.JSON(200, gin.H{
				// "status":   "posted",
				"username": username + "sadf",
				// "password": password,
			})
			// data, _ := ioutil.ReadAll(c.Request.Body)
			// fmt.Printf("ctx.Request.body: %v", string(data))
			// c.Request.ParseForm()
			// for k, v := range c.Request.PostForm {
			// 	fmt.Println("asdk:%v\n", k)
			// 	fmt.Println("asdfasv:%v\n", v)
			// }
			// fmt.Println("sdfasd")
		})
		v1.POST("/submit", func(c *gin.Context) {
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
	e.POST("/ping", func(c *gin.Context) {
		username := c.PostForm("username")
		// password := c.DefaultPostForm("nick", "password") // 此方法可以设置默认值
		fmt.Println(username + "1111")
		c.JSON(200, gin.H{
			// "status":   "posted",
			"username": username + "sadf",
			// "password": password,
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
