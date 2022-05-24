package main

import (
	"CaiNiaoTeam/feedController"
	"CaiNiaoTeam/userController"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", usercontroller.UserInfo)
	apiRouter.POST("/user/register/", usercontroller.Register)
	apiRouter.POST("/user/login/", usercontroller.Login)

	apiRouter.GET("/feed/", feedController.Feed)
}
