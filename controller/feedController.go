package controller

import (
	"CaiNiaoTeam/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

var VideoUrl = common.VideoUrl
var AllVideoList = make([]common.Video, 15, 30) //存放视频列表信息

// RefreshVideoList 刷新视频列表
func RefreshVideoList(list []common.Video) {
	db := common.GetConnection()
	db.Preload("Author").Find(&AllVideoList)
	for i, video := range AllVideoList {
		AllVideoList[i].SetPlayUrl(VideoUrl + "public" + video.PlayUrl)   //拼接视频真正的访问路径，
		AllVideoList[i].SetCoverUrl(VideoUrl + "public" + video.CoverUrl) //如"http://10.34.152.157:8083/"+"public"+"/img/1.jpg"
	}
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	RefreshVideoList(AllVideoList)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: AllVideoList,
		NextTime:  time.Now().Unix(),
	})
}
