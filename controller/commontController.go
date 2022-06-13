package controller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CommentActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

type request_CommentAction struct {
	Token       string `form:"token" json:"token" binding:"required"`
	VideoID     int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType  int64  `form:"action_type" json:"action_type" binding:"required"`
	CommentText string `form:"comment_text" json:"comment_text" binding:"omitempty"`
	CommentID   int64  `form:"comment_id" json:"comment_id" binding:"omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	respond := &request_CommentAction{}

	if err := c.ShouldBind(&respond); err != nil {
		fmt.Println(err.Error())
	}

	user_id := common.TokenParse(respond.Token).(string) //解析token获取用户id

	db := common.GetDB()
	user := new(common.User)
	if db.Where("id = ?", user_id).Find(&user).RowsAffected == 1 { //用户存在
		var comment common.Comment
		switch respond.ActionType {
		case 1: //发布评论

			timeStamp := time.Now().Unix()
			timeLayout := "2006-01-02 15:04:05"
			timeStr := time.Unix(timeStamp, 0).Format(timeLayout)
			comment = common.Comment{UserId: user.Id, VideoId: respond.VideoID, Content: respond.CommentText, CreateDate: timeStr}
			db.Create(&comment)
			fmt.Println(comment)
			db.Preload("User").Find(&comment)

			c.JSON(http.StatusOK, CommentActionResponse{Response: common.Response{StatusCode: 0},
				Comment: comment})
		case 2: //删除评论，
			db.Where("id = ?", respond.CommentID).Delete(&comment)
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		default:
			fmt.Println("action_type错误！")
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

type respond_CommentList struct {
	Video_id string `json:"video_id" form:"video_id"`
	Token    string `json:"token" form:"token"`
}

type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	respond := &respond_CommentList{}

	if err := c.ShouldBind(&respond); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("token:%v\n", respond.Token)
	commentList := getCommentList(respond.Video_id)

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: *commentList,
	})
}

func getCommentList(video_id string) *[]common.Comment {
	db := common.GetDB()

	comments := make([]common.Comment, 20)
	db.Preload("User").Where("video_id = ?", video_id).Find(&comments)

	return &comments
}
