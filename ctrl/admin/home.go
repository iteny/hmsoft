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

type TreeList struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Pid       int         `json:"pid"`
	Sort      int         `json:"sort"`
	Path      string      `json:"path"`
	Component string      `json:"component"`
	Icon      string      `json:"icon"`
	Children  []*TreeList `json:"children"`
}
type DataRes struct {
	Data []*TreeList `json:"data"`
}

func (b *HomeCtrl) GetMenu(c *gin.Context) {
	menu := []sql.Menu{}
	// var menu sql.Menu
	db := common.Sql()
	db.Find(&menu)
	defer db.Close()
	ar := sql.RecursiveMenu(menu, 0, 0)
	fmt.Println(ar)
	c.JSON(200, gin.H{"data": ar})
}
