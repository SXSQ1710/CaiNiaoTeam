package controller

import (
	"CaiNiaoTeam/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: AllVideoList,
	})
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")
	//if len(token) != 0 {
	//	user_id := common.TokenParse(token)
	//
	//}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}
