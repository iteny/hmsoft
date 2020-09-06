package admin

import (
	"fmt"
	"hmsoft/common"
	"hmsoft/sql"

	"github.com/gin-gonic/gin"
)

type SiteCtrl struct {
	common.BaseFunc
}

func SiteCtrlObject() *SiteCtrl {
	return &SiteCtrl{}
}

//获取菜单
func (b *SiteCtrl) GetMenu(c *gin.Context) {
	menu := []sql.Menu{}
	// var menu sql.Menu
	db := common.Sql()
	db.Find(&menu)
	defer db.Close()
	ar := sql.RecursiveMenu(menu, 0, 0)
	fmt.Println(ar)
	c.JSON(200, gin.H{"data": ar})
}

//修改菜单是否显示
func (b *SiteCtrl) ChangeMenuIsshow(c *gin.Context) {
	id := c.PostForm("id")
	isshow := c.PostForm("isshow")
	menu := sql.Menu{}
	db := b.Sql()
	defer db.Close()
	result := db.Model(&menu).Where("id = ?", id).Update("isshow", b.ConvertToInt(isshow)).RowsAffected
	if result == 0 {
		c.JSON(200, gin.H{"status": "faild"})
	} else {
		c.JSON(200, gin.H{"status": "success"})
	}
}
