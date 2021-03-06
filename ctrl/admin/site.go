package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hmsoft/common"
	"hmsoft/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SiteCtrl struct {
	common.BaseFunc
}
type MenuValidationRule struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	Path string `form:"path" json:"path" xml:"path" binding:"required"`
	Icon string `form:"icon" json:"icon" xml:"icon" binding:"required"`
	Pid  *int   `form:"pid" json:"pid" xml:"pid" binding:"required"`
}

func SiteCtrlObject() *SiteCtrl {
	return &SiteCtrl{}
}

//获取菜单
func (b *SiteCtrl) GetMenu(c *gin.Context) {
	menu := []sql.Menu{}
	// var menu sql.Menu
	db := common.Sql()
	db.Order("sort").Find(&menu)
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

//菜单排序
func (b *SiteCtrl) SortMenu(c *gin.Context) {
	sortMenu := make(map[string]string, 0)
	result := c.PostForm("data")
	json.Unmarshal(b.Str2bytes(result), &sortMenu)
	var menuk []string
	for k := range sortMenu {
		menuk = append(menuk, k)
	}
	var bf bytes.Buffer
	bf.WriteString("UPDATE hm_menus SET sort = CASE id ")
	// menu := make(map[string]string, 0)
	for _, v := range menuk {
		// 	mid := vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)
		// 	msort := vali.NumericNoHeadZero(sortMenu[v]) && vali.Length(v, 1, 3)
		// 	switch false {
		// 	case mid:
		// 		c.ResponseJson(11, "Invalid menu id", w, r)
		// 		return
		// 	case msort:
		// 		c.ResponseJson(12, "Invalid sort", w, r)
		// 		return
		// 	default:
		bf.WriteString(fmt.Sprintf("WHEN %v THEN %v ", v, sortMenu[v]))
		// 	}
	}
	// fmt.Println(menu)
	//实现了implode的功能
	var buffer bytes.Buffer
	for _, v := range menuk {
		buffer.WriteString(v)
		buffer.WriteString(",")
	}
	ids := strings.Trim(buffer.String(), ",")
	bf.WriteString(fmt.Sprintf("END WHERE id IN (%v)", ids))
	// fmt.Println(ids)
	fmt.Println(bf.String())
	// tx := c.Sql().MustBegin()
	// tx.MustExec(bf.String())
	// err = tx.Commit()
	// if err != nil {
	// 	c.ResponseJson(4, ""+err.Error(), w, r)
	// } else {
	// 	c.Cache().Del("allmenu")
	// 	c.ResponseJson(1, "", w, r)
	// }
	db := b.Sql()
	defer db.Close()
	res := db.Exec(bf.String())
	if res.RowsAffected == 0 {
		c.JSON(200, gin.H{"status": "faild"})
	} else {
		c.JSON(200, gin.H{"status": "success"})
	}
}

//添加菜单
func (b *SiteCtrl) AddMenu(c *gin.Context) {
	var data MenuValidationRule
	//结构体只和查询参数绑定
	if err := c.ShouldBind(&data); err == nil {
		menu := sql.Menu{Name: data.Name, Path: data.Path, Icon: data.Icon, Pid: *data.Pid}
		// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		db := b.Sql()
		result := db.Create(&menu) // 通过数据的指针来创建
		defer db.Close()
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{"status": "faild"})
		} else {
			c.JSON(200, gin.H{"status": "success"})
		}

	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"res": err.Error()})
	}
}
