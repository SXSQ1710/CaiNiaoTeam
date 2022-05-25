package voideController

import (
	"CaiNiaoTeam/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	db := common.GetConnection()
	db.Preload("Author").Find(&AllVideoList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: *AllVideoList,
	})
}
