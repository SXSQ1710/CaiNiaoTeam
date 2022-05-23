package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	//mysql，dsn
	dsn := "root:13824101958@tcp(localhost:3306)/godemo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	return db
}
