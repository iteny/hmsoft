package admin

import (
	"fmt"
	"hmsoft/common"
	"hmsoft/sql"
	"strconv"

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
	i, _ := strconv.Atoi(isshow)
	menu := sql.Menu{}
	db := b.Sql()
	defer db.Close()
	// tx := db.Begin()
	db.Debug().Model(&menu).Where("id = ?", id).Update("isshow", i)
	// if result.RowsAffected == 0 {
	// 	//如果更新库存操作，返回影响行数为0，说明没有库存了，结束下单流程
	// 	//这里回滚作用不大，因为前面没成功执行什么数据库更新操作，也没什么数据需要回滚。
	// 	//这里就是举个例子，事务中可以执行多个sql语句，错误了可以回滚事务
	// 	tx.Rollback()
	// 	return
	// }
	// tx.Commit()
}
