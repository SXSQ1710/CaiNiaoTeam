package controller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.View_video_favorites `json:"video_list,omitempty"`
	NextTime  int64                         `json:"next_time,omitempty"`
}

var VideoUrl = common.VideoUrl
var AllVideoList = make([]common.View_video_favorites, 15, 30) //存放视频列表信息

// RefreshVideoList 刷新视频列表 TODO 待修饰
func RefreshVideoList(list []common.View_video_favorites, user_id string) []common.View_video_favorites {
	db := common.GetDB()
	db.Preload("Author").Order("id desc").Find(&list)
	favorite := make([]common.Favorite, len(list))
	db.Find(&favorite)
	AllVideoMap := make(map[int64]*common.View_video_favorites, len(list))
	for i, video := range list {
		list[i].PlayUrl = VideoUrl + "public" + video.PlayUrl   //拼接视频真正的访问路径，
		list[i].CoverUrl = VideoUrl + "public" + video.CoverUrl //如"http://10.34.152.157:8083/"+"public"+"/img/1.jpg"
		AllVideoMap[video.Id] = &list[i]
	}
	for _, f := range favorite {

		if f.User_id == user_id {
			if video_id_Int, err := strconv.ParseInt(f.Video_id, 10, 64); err != nil {
				fmt.Printf("Video_id转换失败！,err:%v\n", err)
			} else {
				AllVideoMap[video_id_Int].IsFavorite = true
			}
		}
	}

	return list
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	id := common.TokenParse(token).(string)
	AllVideoList = RefreshVideoList(AllVideoList, id)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: AllVideoList,
		NextTime:  time.Now().Unix(),
	})
}

/**
* 响应“/feed/”
* 上面部分
---------------------------------------------------分界线----------------------------------------------------------------
* 下面部分
* 测试
**/

func SetUrlA(list []common.View_video_favorites) []common.View_video_favorites {
	for i, video := range list {
		list[i].PlayUrl = VideoUrl + "public" + video.PlayUrl
		list[i].CoverUrl = VideoUrl + "public" + video.CoverUrl
		//list[i].SetPlayUrl(VideoUrl + "public" + video.PlayUrl)   //拼接视频真正的访问路径，
		//list[i].SetCoverUrl(VideoUrl + "public" + video.CoverUrl) //如"http://10.34.152.157:8083/"+"public"+"/img/1.jpg"
	}
	return list

}

func SetUrlB(list []common.View_video_favorites) []common.View_video_favorites {
	for i, video := range list {
		//list[i].PlayUrl = VideoUrl + "public" + video.PlayUrl
		//list[i].CoverUrl = VideoUrl + "public" + video.CoverUrl
		list[i].SetPlayUrl(VideoUrl + "public" + video.PlayUrl)   //拼接视频真正的访问路径，
		list[i].SetCoverUrl(VideoUrl + "public" + video.CoverUrl) //如"http://10.34.152.157:8083/"+"public"+"/img/1.jpg"
	}
	return list
}

func SetUrlTestA(list []common.View_video_favorites) {
	AllVideoList = SetUrlA(list)
}

func SetUrlTestB(list []common.View_video_favorites) {
	AllVideoList = SetUrlB(list)
}
