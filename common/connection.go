package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func GetDB() *gorm.DB {
	fmt.Println("获取数据库连接！")
	return _db
}

func init() {

	dsn := Dsn
	var err error

	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("数据库连接失败:%v\n", err)
	}
	fmt.Println("获取数据库连接池成功！")
}
