package usercontroller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type UserLoginResponse struct {
	common.Response        //统一响应结构
	UserId          int64  `json:"user_id,omitempty"`
	Token           string `json:"token"` //用户身份鉴权认证:https://zhuanlan.zhihu.com/p/433197184
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

var CreateUserTable sync.Once

func Register(c *gin.Context) {
	//获取用户输入的username和password
	username := c.Query("username")
	password := c.Query("password")

	token := username + password //用户身份鉴权认证
	fmt.Printf("用户鉴权：%v\n", token)
	CreateUserTable.Do(fn_creatUserTable) //创建user表，只运行一次

	user := new(common.User)
	db := common.GetConnection()
	out := db.Where("token = ?", token).Find(&user)
	//fmt.Printf("返回结果数目:%v\n", out.RowsAffected)

	if out.RowsAffected == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户已经存在！"},
		})
	} else {
		db.Create(&common.User{Token: token, Name: username})
		db.Where("token = ?", token).Find(&user)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    username + password,
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user := new(common.User)
	if common.GetConnection().Where("token = ?", token).Find(&user).RowsAffected == 1 {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户不存在！"},
		})
	}
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
