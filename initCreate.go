package main

import (
	"CaiNiaoTeam/common"
	"fmt"
	"sync"
)

var CreateUserTable sync.Once

func initCreatTable() {
	CreateUserTable.Do(fn_creatUserTable) //创建user表，只运行一次
}

func fn_creatUserTable() {
	db := common.GetConnection()

	if !(db.Migrator().HasTable("userinfo")) {
		if err := db.Table("users").Migrator().CreateTable(&common.User{}); err != nil {
			fmt.Println("fn_creatUserTable:" + err.Error())
		}
	}
	db.Exec("alter table users AUTO_INCREMENT = 10000")

	fmt.Println("运行fn_creatUserTable")
}
