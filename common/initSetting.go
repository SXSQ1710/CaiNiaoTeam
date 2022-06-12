package common

import (
	"fmt"
	"sync"
)

var Dsn = "root:13824101958@tcp(localhost:3306)/godemo1" //你自己数据库dsn
var mySigningKey = []byte("qwacdfafaefa.")               //你自己token加密解密的秘钥，可随便设置
var SetTime int64 = 12 * 60 * 60 * 30                    //token过期时间
var VideoUrl = "http://10.34.152.157:8083/"              //填写你本地资源的访问入口,我用的是nginx开启本地资源入口
//var VideoUrl = "http://10.31.46.13:8083/"                            //填写你本地资源的访问入口,我用的是nginx开启本地资源入口
var LocalUrl = "D:\\GolandProjects\\github.com\\CaiNiaoTeam\\public" //填写你视频和视频封面的本地地址

//这是本地feature分支的测试
func InitCreatTable() {
	var CreateUserTable sync.Once
	var CreateVideoTable sync.Once
	var CreatFavoriteTable sync.Once
	var CreatView_video_favorites sync.Once
	var CreatCommentTable sync.Once
	var CreatRelationTable sync.Once
	var CreatView_user_relations sync.Once
	CreateUserTable.Do(creatUserTable)                      //创建user表，只运行一次
	CreateVideoTable.Do(creatVideoTable)                    //创建video表，只运行一次
	CreatFavoriteTable.Do(creatFavoriteTable)               //创建Favorite表，只运行一次
	CreatCommentTable.Do(creatCommentTable)                 //创建comments视图，只运行一次
	CreatRelationTable.Do(creatRelationTable)               //创建relation表，只运行一次
	CreatView_video_favorites.Do(creatView_video_favorites) //创建view_video_favorites视图，只运行一次
	CreatView_user_relations.Do(creatView_user_relations)   //创建view_user_relations视图，只运行一次
	fmt.Println("-----------------------------")
	fmt.Println("运行InitCreatTable")
	fmt.Println("-----------------------------")
}

func creatUserTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("users")) {
		if err := db.Table("users").Migrator().CreateTable(&User{}); err != nil {
			fmt.Println("creatUserTable:" + err.Error())
		}
	}
	db.Exec("alter table users AUTO_INCREMENT = 10000") //id字段从10000开始自动自增

}

func creatVideoTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("videos")) {
		if err := db.Table("videos").Migrator().CreateTable(&Video{}); err != nil {
			fmt.Println("creatVideoTable:" + err.Error())
		}
	}

}

func creatFavoriteTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("favorites")) {
		if err := db.Table("favorites").Migrator().CreateTable(&Favorite{}); err != nil {
			fmt.Println("creatFavoriteTable:" + err.Error())
		}
	}
}

func creatRelationTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("relations")) {
		if err := db.Table("relations").Migrator().CreateTable(&Relation{}); err != nil {
			fmt.Println("creatRelationTable:" + err.Error())
		}
	}
}

func creatCommentTable() {
	db := GetDB()

	if !(db.Migrator().HasTable("comments")) {
		if err := db.Table("comments").Migrator().CreateTable(&Comment{}); err != nil {
			fmt.Println("creatCommentTable:" + err.Error())
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
		db.Create(&Video{AuthorId: user.Id, Title: "初始视频2", PlayUrl: "/video/10000_2.mp4", CoverUrl: "/img/10000_2.jpg"})
		fmt.Println("-----------------------------")
		fmt.Println("运行AddInitInfo")
		fmt.Println("-----------------------------")
	}
}

func creatView_video_favorites() {
	db := GetDB()
	sql := "CREATE VIEW view_video_favorites(id,author_id,play_url,cover_url,favorite_count,comment_count,is_favorite,title)\nAS\nSELECT v.id,v.author_id,v.play_url,v.cover_url,count(DISTINCT f.user_id),count(DISTINCT c.content),v.is_favorite,v.title\nFROM (`videos` v) LEFT JOIN (`favorites` f) ON v.id = f.video_id  LEFT JOIN (`comments` c) ON v.id = c.video_id\nGROUP BY v.id "

	if !(db.Migrator().HasTable("view_video_favorites")) {
		if err := db.Exec(sql).Error; err != nil {
			fmt.Println("creatView_video_favorites:" + err.Error())
		}
	}

}

func creatView_user_relations() {
	db := GetDB()
	sql := "CREATE VIEW view_user_relations(id,id_pass,name,follow_count,follower_count,is_follow)\nAS\nSELECT u.id,u.id_pass,u.name,count(DISTINCT r.follow_user_id),\n (SELECT COUNT(DISTINCT user_id)\n FROM `relations`\n WHERE follow_user_id = u.id)\n ,u.is_follow\nFROM (`users` u) LEFT JOIN (`relations` r) ON u.id = r.user_id \nGROUP BY u.id "

	if !(db.Migrator().HasTable("view_user_relations")) {
		if err := db.Exec(sql).Error; err != nil {
			fmt.Println("creatView_user_relations:" + err.Error())
		}
	}

}
