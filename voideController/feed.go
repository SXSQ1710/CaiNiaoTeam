package voideController

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

var VideoUrl = "http://10.34.152.157:8083/" //填写你本地资源的访问入口

var AllVideoList = make([]common.Video, 30)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	db := common.GetConnection()
	db.Preload("Author").Find(&AllVideoList)
	for i, video := range AllVideoList {
		AllVideoList[i].SetPlayUrl(VideoUrl + "public" + video.PlayUrl)
		AllVideoList[i].SetCoverUrl(VideoUrl + "public" + video.CoverUrl)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: AllVideoList,
		NextTime:  time.Now().Unix(),
	})
}

func setVideoAuthor(video common.Video, user *common.User) {

}
