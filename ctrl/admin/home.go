package admin

import (
	"fmt"
	"hmsoft/common"
	"hmsoft/sql"

	"github.com/gin-gonic/gin"
)

type HomeCtrl struct {
	common.BaseFunc
}

func HomeCtrlObject() *HomeCtrl {
	return &HomeCtrl{}
}

func (b *HomeCtrl) GetMenu(c *gin.Context) {
	menu := []sql.Menu{}
	// var menu sql.Menu
	db := common.Sql()
	db.Order("sort").Find(&menu)
	defer db.Close()
	ar := sql.RecursiveMenu(menu, 0, 0)
	fmt.Println(ar)
	c.JSON(200, gin.H{"data": ar})
}
