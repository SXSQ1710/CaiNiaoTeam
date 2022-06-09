package controller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	select_userId := c.Query("user_id")          //这里的user_id是查询用户的id
	user_id := common.TokenParse(token).(string) //这里是用户自己的id

	userAllLikeVideo := getUserLikeVideoList(AllVideoList, select_userId, user_id)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: userAllLikeVideo,
	})
}

// 获取用户所有喜欢的视频的视频列表
func getUserLikeVideoList(list []common.View_video_favorites, select_userId string, user_id string) []common.View_video_favorites {
	sql := "SELECT v.id,v.author_id,v.play_url,v.cover_url,v.favorite_count,v.comment_count,v.is_favorite,v.title FROM view_video_favorites v,favorites f WHERE f.video_id = v.id AND f.user_id = ?"
	db := common.GetDB()

	if err := db.Preload("Author").Raw(sql, select_userId).Scan(&list).Error; err != nil {
		log.Println(err.Error())
	}

	//db.Preload("Author").Find(&list)
	favorite := make([]common.Favorite, len(list))
	db.Where("user_id = ?", user_id).Find(&favorite)
	AllVideoMap := make(map[int64]*common.View_video_favorites, len(list))
	for i, video := range list {
		list[i].PlayUrl = VideoUrl + "public" + video.PlayUrl   //拼接视频真正的访问路径，
		list[i].CoverUrl = VideoUrl + "public" + video.CoverUrl //如"http://10.34.152.157:8083/"+"public"+"/img/1.jpg"
		AllVideoMap[video.Id] = &list[i]
	}

	fmt.Println(list)
	fmt.Println(favorite)

	//这一步是将查看用户的所有喜欢视频中存在和用户一样喜欢的视频的IsFavorite标为true
	//但结果发现他app里根本没做这项功能，我像一个小丑一样弄了好久(＠_＠)
	for _, f := range favorite { //循环用户的喜欢视频列表
		video_id_Int, err := strconv.ParseInt(f.Video_id, 10, 64) //将视频id转换为int64类型
		if err != nil {
			fmt.Printf("Video_id转换失败！,err:%v\n", err)
		}
		if _, exist := AllVideoMap[video_id_Int]; exist { //判断在查看用户的喜欢视频中是否存在用户一样喜欢的视频
			fmt.Println(video_id_Int)
			AllVideoMap[video_id_Int].IsFavorite = true
		}
	}

	return list
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	//user_id := c.Query("user_id")

	if len(token) != 0 {
		db := common.GetDB()
		id := common.TokenParse(token).(string)
		favorite := common.Favorite{User_id: id, Video_id: video_id}

		if action_type == "1" {
			db.Create(&favorite)
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		} else if action_type == "2" {
			fmt.Println("取消点赞")
			db.Delete(&favorite)
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		} else {
			fmt.Println("action_type错误！")
		}

	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户token错误！"},
		})
	}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}
