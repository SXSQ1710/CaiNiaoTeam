package common

import (
	"fmt"
	"sync"
)

var Dsn = "root:13824101958@tcp(localhost:3306)/godemo1" //你自己数据库dsn

var mySigningKey = []byte("qwacdfafaefa.") //你自己token加密解密的秘钥，可随便设置

var SetTime int64 = 2 * 60 * 60 //token过期时间

var VideoUrl = "http://10.34.152.157:8083/" //填写你本地资源的访问入口,我用的是nginx开启本地资源入口

func InitCreatTable() {
	var CreateUserTable sync.Once
	var CreateVideoTable sync.Once
	CreateUserTable.Do(fn_creatUserTable)   //创建user表，只运行一次
	CreateVideoTable.Do(fn_creatVideoTable) //创建video表，只运行一次
}

func fn_creatUserTable() {
	db := GetConnection()

	if !(db.Migrator().HasTable("users")) {
		if err := db.Table("users").Migrator().CreateTable(&User{}); err != nil {
			fmt.Println("fn_creatUserTable:" + err.Error())
		}
	}
	db.Exec("alter table users AUTO_INCREMENT = 10000") //id字段从10000开始自动自增

	fmt.Println("运行fn_creatUserTable")
}

func fn_creatVideoTable() {
	db := GetConnection()

	if !(db.Migrator().HasTable("videos")) {
		if err := db.Table("videos").Migrator().CreateTable(&Video{}); err != nil {
			fmt.Println("fn_creatVideoTable:" + err.Error())
		}
	}

	fmt.Println("运行fn_creatVideoTable")
}
