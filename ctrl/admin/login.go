package admin

import (
	"fmt"
	"hmsoft/common"
	"hmsoft/sql"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginCtrl struct {
	common.BaseFunc
}
type Login struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginCtrlObject() *LoginCtrl {
	return &LoginCtrl{}
}

//用户登录
func (b *LoginCtrl) LoginSubmit(c *gin.Context) {
	var data Login
	//结构体只和查询参数绑定
	if err := c.ShouldBind(&data); err == nil {
		// var ss sql.User
		userStruct := sql.User{Username: data.Username, Password: b.Sha1PlusMd5(data.Password)}
		// fmt.Println(user)
		db := b.Sql()
		e := db.Where(&userStruct).First(&userStruct)
		defer db.Close()
		// fmt.Println(&ss)
		if e.Error != nil {
			c.JSON(200, gin.H{"status": "faild", "info": e.Error.Error()})
			fmt.Println(e.Error)
		} else {
			if userStruct.Id != 0 {
				if userStruct.Status == 1 {
					session := sessions.Default(c)
					if session.Get("uid") != userStruct.Id {
						session.Set("uid", userStruct.Id)
						session.Set("user", data.Username)
						session.Save()
					}
					tx := db.Begin()
					loginLogStuct := sql.LoginLog{Username: data.Username, Time: time.Now().Unix(), Ip: c.ClientIP(), Useragent: c.GetHeader("User-Agent"), Uid: userStruct.Id}
					if err := tx.Create(&loginLogStuct).Error; err != nil {
						tx.Rollback()
						c.JSON(200, gin.H{"res": err.Error()})
					} else {
						tx.Commit()
					}
					c.JSON(200, gin.H{"status": "success"})
				}
			} else {
				c.JSON(200, gin.H{"status": "faild"})
			}
		}
	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"res": err.Error()})
	}
}

//登录的状态
func (b *LoginCtrl) LoginStatus(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Println(session.Get("uid"))
	if session.Get("uid") != nil {
		c.JSON(200, gin.H{"status": "success", "info": session.Get("user")})
	} else {
		c.JSON(200, gin.H{"status": "faild"})
	}
}
