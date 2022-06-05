package controller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserResponse 用户信息
type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")

	id := common.TokenParse(token)

	if id == user_id {
		user := new(common.User)
		common.GetConnection().Where("id = ?", user_id).Find(&user)
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户信息错误！"},
		})
	}
}

/**
* 用户信息获取对于应url为"/douyin/user/"的请求
* 上面部分
---------------------------------------------------分界线----------------------------------------------------------------
* 下面部分
* 用户登录对于应url为"/douyin/user/login/"的请求
**/

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	idPass := username + password //数据库中可以唯一标识用户
	user := new(common.User)
	out := common.GetConnection().Where("id_pass = ?", idPass).Find(&user).RowsAffected

	id := fmt.Sprintf("%d", user.Id)

	token := common.TokenEncrypt(id) //用jwt将user_id作为token进行加密

	if out == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

/**
* 用户登录对于应url为"/douyin/user/login/"的请求
* 上面部分
---------------------------------------------------分界线----------------------------------------------------------------
* 下面部分
* 用户注册对于应url为"/douyin/user/register/"的请求
*/

// UserLoginResponse 用户登录响应结构体
type UserLoginResponse struct {
	common.Response        //统一响应结构
	UserId          int64  `json:"user_id,omitempty"`
	Token           string `json:"token"` //用户身份鉴权认证:https://zhuanlan.zhihu.com/p/433197184
}

func Register(c *gin.Context) {
	//获取用户输入的username和password
	username := c.Query("username")
	password := c.Query("password")

	idPass := common.BuilderString(username, password) //用户身份鉴权认证

	user := new(common.User)
	db := common.GetConnection()
	out := db.Where("id_pass = ?", idPass).Find(&user) //从数据库中查询用户是否存在，直接查询idPass是否存在
	//fmt.Printf("返回结果数目:%v\n", out.RowsAffected)

	if out.RowsAffected == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户已经存在！"},
		})
	} else {
		db.Create(&common.User{IdPass: idPass, Name: username})
		db.Where("id_pass = ?", idPass).Find(&user)
		id := fmt.Sprintf("%d", user.Id)
		token := common.TokenEncrypt(id) //用jwt将user_id作为token进行加密
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}
