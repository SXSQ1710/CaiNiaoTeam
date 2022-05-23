package usercontroller

import (
	"CaiNiaoTeam/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user := new(common.User)
	if common.GetConnection().Where("token = ?", token).Find(&user).RowsAffected == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    username + password,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
