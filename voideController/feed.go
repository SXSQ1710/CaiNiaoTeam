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

var AllVideoList = new([]common.Video)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	AllVideoList := new([]common.Video)
	db := common.GetConnection()
	db.Preload("Author").Find(&AllVideoList)
	//for _, video := range *AllVideoList {
	//	db := common.GetConnection()
	//	user := new(common.User)
	//	db.Where("id = ?", video.AuthorId).Find(&user)
	//	fmt.Println(user)
	//}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: *AllVideoList,
		NextTime:  time.Now().Unix(),
	})
}

func setVideoAuthor(video common.Video, user *common.User) {

}
