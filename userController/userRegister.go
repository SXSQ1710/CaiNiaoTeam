package usercontroller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Register(c *gin.Context) {
	//获取用户输入的username和password
	username := c.Query("username")
	password := c.Query("password")

	token := username + password //用户身份鉴权认证
	fmt.Printf("用户鉴权：%v\n", token)

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
