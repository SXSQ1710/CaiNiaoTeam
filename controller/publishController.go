package controller

import (
	"CaiNiaoTeam/common"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	//token := c.Query("token")    //这里发现token没啥用
	userId := c.Query("user_id") //这里的user_id是查询用户的id

	AllVideoList = getUserVideoList(AllVideoList, userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: AllVideoList,
	})
	//if len(token) != 0 {
	//	id := common.TokenParse(token)
	//	if id == userId { //身份验证通过
	//		AllVideoList = getUserVideoList(AllVideoList, userId)
	//		c.JSON(http.StatusOK, VideoListResponse{
	//			Response: common.Response{
	//				StatusCode: 0,
	//			},
	//			VideoList: AllVideoList,
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, UserResponse{
	//			Response: common.Response{StatusCode: 1, StatusMsg: "用户信息错误！"},
	//		})
	//	}
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: common.Response{StatusCode: 1, StatusMsg: "用户信息错误！"},
	//	})
	//}
}

// 获取用户所有发布视频的视频列表
func getUserVideoList(list []common.Video, user_id string) []common.Video {

	db := common.GetDB()
	db.Preload("Author").Where("author_id = ?", user_id).Find(&list)
	for i, video := range list {
		list[i].PlayUrl = VideoUrl + "public" + video.PlayUrl   //拼接视频真正的访问路径，
		list[i].CoverUrl = VideoUrl + "public" + video.CoverUrl //如"http://10.34.152.157:8083/"+"public"+"/img/1.jpg"
	}
	return list
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
	db := common.GetDB()
	out := db.Where("id = ?", parse).Find(&user).RowsAffected

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
	unixTime := time.Now().Unix()

	videoName := fmt.Sprintf("%d_%d_%s", user.Id, unixTime, filename) //拼接视频名称:用户id+上传时间+文件名
	saveFile := filepath.Join("./public/video/", videoName)           //存储视频到本机上
	coverName := fmt.Sprintf("%d_%d.jpeg", user.Id, unixTime)         //拼接视频封面名称:用户id+上传时间+".jpeg"
	defer creatCover(videoName, coverName)                            //使用ffmpeg截取视频生成封面
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	////写入数据库视频信息
	playUrl := common.BuilderString("/video/", videoName)
	coverUrl := common.BuilderString("/img/", coverName)
	db.Create(&common.Video{AuthorId: user.Id, Title: title, PlayUrl: playUrl, CoverUrl: coverUrl})

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  filename + " uploaded successfully",
	})
}

//使用ffmpeg截取视频生成封面
func creatCover(videoName string, coverName string) {
	videoLocalUrl := common.LocalUrl + "\\video\\" + videoName
	coverLocalUrl := common.LocalUrl + "\\img\\" + coverName
	fmt.Println(videoLocalUrl)
	fmt.Println(coverLocalUrl)
	cmdArguments := []string{"-i", videoLocalUrl, "-ss", "00:00:01", "-t", "1", "-r", "1", "-q:v", "2", "-f", "image2", coverLocalUrl}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("command output: %q", out.String())
}
