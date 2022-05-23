package usercontroller

import (
	"CaiNiaoTeam/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
