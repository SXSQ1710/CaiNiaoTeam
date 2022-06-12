package middleWare

import (
	"CaiNiaoTeam/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type unAuthenticationRespond struct {
	Token   string `json:"token" form:"token"`
	User_id string `json:"user_id" form:"user_id" binding:"omitempty"`
}

func Authentication(c *gin.Context) {
	respond := &unAuthenticationRespond{}
	if err := c.ShouldBind(&respond); err != nil {
		fmt.Println(err.Error())
	}

	if respond.User_id != "" && respond.Token != "" {
		parse_id := common.TokenParse(respond.Token)
		if parse_id == respond.User_id {
			start := time.Now().UnixNano()
			c.Next()
			end := time.Now().UnixNano()
			fmt.Printf("运行时间：%d", end-start)
		} else {
			fmt.Println("请求拒绝1")
			c.Abort()
		}
	} else if respond.Token != "" {
		user_id := common.TokenParse(respond.Token).(string)
		db := common.GetDB()
		if db.Where("id = ?", user_id).Find(&common.User{}).RowsAffected == 1 {
			c.Next()
		} else {
			fmt.Println("请求拒绝2")
			c.Abort()
		}
	} else {
		fmt.Println("请求拒绝3")
		c.Abort()
	}
}
