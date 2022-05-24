package voideController

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
		VideoList: *AllVideoList,
	})
}
