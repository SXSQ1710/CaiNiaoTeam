package main

import (
	"CaiNiaoTeam/common"
	"fmt"
	"sync"
)

var CreateUserTable sync.Once
var CreateVideoTable sync.Once

func initCreatTable() {
	CreateUserTable.Do(fn_creatUserTable)   //创建user表，只运行一次
	CreateVideoTable.Do(fn_creatVideoTable) //创建video表，只运行一次
}

func fn_creatUserTable() {
	db := common.GetConnection()

	if !(db.Migrator().HasTable("users")) {
		if err := db.Table("users").Migrator().CreateTable(&common.User{}); err != nil {
			fmt.Println("fn_creatUserTable:" + err.Error())
		}
	}
	db.Exec("alter table users AUTO_INCREMENT = 10000")

	fmt.Println("运行fn_creatUserTable")
}

func fn_creatVideoTable() {
	db := common.GetConnection()

	if !(db.Migrator().HasTable("videos")) {
		if err := db.Table("videos").Migrator().CreateTable(&common.Video{}); err != nil {
			fmt.Println("fn_creatVideoTable:" + err.Error())
		}
	}

	fmt.Println("运行fn_creatVideoTable")
}
