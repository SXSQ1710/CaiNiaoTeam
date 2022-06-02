package controller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")

	id := common.TokenParse(token)
	if id == user_id {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			VideoList: AllVideoList,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户信息错误！"},
		})
	}
}

/**
* 响应“/publish/list/”
* 上面部分
---------------------------------------------------分界线----------------------------------------------------------------
* 下面部分
* 响应“/publish/action/”
**/

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	parse := common.TokenParse(token)

	user := new(common.User)
	out := common.GetConnection().Where("id = ?", parse).Find(&user).RowsAffected

	if out == 0 {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/video/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	////写入数据库视频信息
	playUrl := common.BuilderString("/video/", finalName)
	db := common.GetConnection()
	db.Create(&common.Video{AuthorId: user.Id, Title: title, PlayUrl: playUrl})

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}
