package main

import (
	"CaiNiaoTeam/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
}
