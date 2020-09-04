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
