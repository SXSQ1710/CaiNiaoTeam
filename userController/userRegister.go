package usercontroller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	common.Response        //统一响应结构
	UserId          int64  `json:"user_id,omitempty"`
	Token           string `json:"token"` //用户身份鉴权认证:https://zhuanlan.zhihu.com/p/433197184
}

var CreateUserTable sync.Once

func Register(c *gin.Context) {
	//获取用户输入的username和password
	username := c.Query("username")
	password := c.Query("password")

	token := username + password //用户身份鉴权认证
	fmt.Printf("用户鉴权：%v", token)

	CreateUserTable.Do(fn_creatUserTable) //创建user表，只运行一次

	// 	// 获取所有匹配的记录
	// db.Where("name = ?", "jinzhu").Find(&users)
	// //// SELECT * FROM users WHERE name = 'jinzhu';
	user := new(Userinfo)
	db := common.GetConnection()
	windows := db.Where("token = ?", token).Find(&user)
	fmt.Println(windows)
	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	// 	})
	// } else {
	// 	atomic.AddInt64(&userIdSequence, 1)
	// 	newUser := User{
	// 		Id:   userIdSequence,
	// 		Name: username,
	// 	}
	// 	usersLoginInfo[token] = newUser
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 0},
	// 		UserId:   userIdSequence,
	// 		Token:    username + password,
	// 	})
	// }
}

//User模型
type Userinfo struct {
	Username int `gorm:"primary_key"`
	Password int `gorm:"size:32"`
	Token    int `gorm:"size:64"`
}

func fn_creatUserTable() {
	db := common.GetConnection()

	if !(db.Migrator().HasTable("userinfo")) {
		if err := db.Table("userinfos").Migrator().CreateTable(&Userinfo{}); err != nil {
			fmt.Println("fn_creatUserTable:" + err.Error())
		}
	}
	fmt.Println("运行fn_creatUserTable")
}
