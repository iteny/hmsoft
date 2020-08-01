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
	return DB
}
func (c *BaseFunc) Sql() *gorm.DB {
	return DB
}
func init() {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "hm_" + defaultTableName
	// }
	var err error
	DB, err = gorm.Open("mysql", "root:iteny@/hmsoft?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// defer DB.Close()
	// var ss sql.User
	// DB.Table("user").First(&ss)
	// fmt.Println(&ss)
	// fmt.Println(err)
}
