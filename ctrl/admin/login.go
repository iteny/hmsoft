package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCtrl struct {
}
type Login struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginCtrlObject() *LoginCtrl {
	return &LoginCtrl{}
}
func (t *LoginCtrl) LoginSubmit(c *gin.Context) {
	// username := c.PostForm("username")
	// // password := c.DefaultPostForm("nick", "password") // 此方法可以设置默认值
	// fmt.Println(username + "1111")
	// c.JSON(200, gin.H{
	// 	// "status":   "posted",
	// 	"username": username + "sadf",
	// 	// "password": password,
	// })
	var data Login
	//结构体只和查询参数绑定
	if err := c.ShouldBind(&data); err == nil {
		fmt.Println(data.Username)
		fmt.Println(data.Password)
		c.JSON(http.StatusOK, gin.H{"res": "success"})
	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"res": err.Error()})
	}
}
