package usercontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	fmt.Println(token)
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
