package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BaseFunc struct {
	// LogCtrl
}

var DB *gorm.DB

func Sql() *gorm.DB {
	DB, err := gorm.Open("mysql", "root:iteny@/hmsoft?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}

//打开一个数据库连接需要手动关闭
func (c *BaseFunc) Sql() *gorm.DB {
	DB, err := gorm.Open("mysql", "root:iteny@/hmsoft?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}
func init() {
	//设置数据库默认表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "hm_" + defaultTableName
	}
}
