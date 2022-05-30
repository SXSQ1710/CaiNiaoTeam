package common

import (
	"CaiNiaoTeam/initSetting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	//mysql，dsn
	dsn := initSetting.Dsn

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	return db
}
