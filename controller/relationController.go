package controller

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type request_RelationAction struct {
	Token       string `form:"token" json:"token" binding:"required"`
	User_id     int64  `form:"user_id" json:"user_id" binding:"omitempty" `
	To_user_id  string `form:"to_user_id" json:"to_user_id" binding:"required"`
	Action_type int64  `form:"action_type" json:"action_type" binding:"required"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	respond := &request_RelationAction{}
	if err := c.ShouldBind(&respond); err != nil {
		fmt.Println(err.Error())
	}

	user_id := common.TokenParse(respond.Token).(string)

	db := common.GetDB()
	relation := common.Relation{User_id: user_id, Follow_user_id: respond.To_user_id}
	if respond.Action_type == 1 {
		if err := addRelation(&relation, db); err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Relation error!"})
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		}
	} else if respond.Action_type == 2 {
		if err := deleteRelation(&relation, db); err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "Relation error!"})
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		}
	}
}

func addRelation(relation *common.Relation, db *gorm.DB) error {
	return db.Create(&relation).Error
}
func deleteRelation(relation *common.Relation, db *gorm.DB) error {
	return db.Delete(&relation).Error
}

type request_FollowList struct {
	Token          string `form:"token" json:"token" binding:"omitempty"`
	Select_user_id int64  `form:"user_id" json:"user_id" binding:"omitempty" ` //这里是查询的user_id
}

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	respond := &request_FollowList{}
	if err := c.ShouldBind(&respond); err != nil {
		fmt.Println(err.Error())
	}

	db := common.GetDB()

	if follow := getAllUserFollow(respond.Select_user_id, db); follow == nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "FollowList error!"})
	} else {
		fmt.Println(follow)
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			UserList: *follow,
		})
	}

}

func getAllUserFollow(select_user_id int64, db *gorm.DB) *[]common.User {
	relations := []common.Relation{}
	followList := []common.User{}
	if err := db.Preload("Follow_user").Where("user_id = ?", select_user_id).Find(&relations).Error; err != nil {
		return nil
	} else {
		for i := 0; i < len(relations); i++ {
			followList = append(followList, relations[i].Follow_user)
			followList[i] = relations[i].Follow_user
		}
		return &followList
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {

	respond := &request_FollowList{}
	if err := c.ShouldBind(&respond); err != nil {
		fmt.Println(err.Error())
	}

	db := common.GetDB()

	if follow := getAllFollowUser(respond.Select_user_id, db); follow == nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "FollowList error!"})
	} else {
		fmt.Println(follow)
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			UserList: *follow,
		})
	}
}

func getAllFollowUser(select_user_id int64, db *gorm.DB) *[]common.User {
	relations := []common.Relation{}
	followList := []common.User{}
	if err := db.Preload("Follower_user").Where("Follow_user_id = ?", select_user_id).Find(&relations).Error; err != nil {
		return nil
	} else {
		for i := 0; i < len(relations); i++ {
			followList = append(followList, relations[i].Follower_user)
			followList[i] = relations[i].Follower_user
		}
		return &followList
	}
}
