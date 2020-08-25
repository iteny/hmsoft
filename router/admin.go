//后端路由
package router

import (
	"fmt"
	"hmsoft/ctrl/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouter(e *gin.Engine) {
	login := admin.LoginCtrlObject()
	v1 := e.Group("/admin")
	{
		v1.POST("/login", login.LoginSubmit) //登录api
		v1.POST("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})
		v1.POST("/loginstatus", login.LoginStatus) //用户登录状态验证
		v1.POST("/loginout", login.LoginOut)       //退出登录
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
