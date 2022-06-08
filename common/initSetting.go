package common

import (
	"fmt"
	"sync"
)

var Dsn = "root:13824101958@tcp(localhost:3306)/godemo1"             //你自己数据库dsn
var mySigningKey = []byte("qwacdfafaefa.")                           //你自己token加密解密的秘钥，可随便设置
var SetTime int64 = 2 * 60 * 60                                      //token过期时间
var VideoUrl = "http://10.34.152.157:8083/"                          //填写你本地资源的访问入口,我用的是nginx开启本地资源入口
var LocalUrl = "D:\\GolandProjects\\github.com\\CaiNiaoTeam\\public" //填写你视频和视频封面的本地地址
//这是本地feature分支的测试
func InitCreatTable() {
	var CreateUserTable sync.Once
	var CreateVideoTable sync.Once
	CreateUserTable.Do(creatUserTable)   //创建user表，只运行一次
	CreateVideoTable.Do(creatVideoTable) //创建video表，只运行一次
	fmt.Println("-----------------------------")
	fmt.Println("运行InitCreatTable")
	fmt.Println("-----------------------------")
}

func creatUserTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("users")) {
		if err := db.Table("users").Migrator().CreateTable(&User{}); err != nil {
			fmt.Println("fn_creatUserTable:" + err.Error())
		}
	}
	db.Exec("alter table users AUTO_INCREMENT = 10000") //id字段从10000开始自动自增

}

func creatVideoTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("videos")) {
		if err := db.Table("videos").Migrator().CreateTable(&Video{}); err != nil {
			fmt.Println("fn_creatVideoTable:" + err.Error())
		}
	}

}

func AddInitInfo() {
	db := GetDB()

	user := new(User)
	if db.Where("id_pass = ?", "SXSQ123456").Find(&user).RowsAffected == 0 {
		db.Create(&User{IdPass: "SXSQ123456", Name: "SXSQ"}) //添加初始用户
		db.Where("id_pass = ?", "SXSQ123456").Find(&user)
		db.Create(&Video{AuthorId: user.Id, Title: "初始视频1", PlayUrl: "/video/10000_1.mp4", CoverUrl: "/img/10000_1.jpg"})
		db.Create(&Video{AuthorId: user.Id, Title: "初始视频2", PlayUrl: "/video/10000_2.mp4", CoverUrl: "/img/10000_1.jpg"})
		fmt.Println("-----------------------------")
		fmt.Println("运行AddInitInfo")
		fmt.Println("-----------------------------")
	}
}
