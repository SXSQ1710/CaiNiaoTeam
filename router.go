package main

import (
	"CaiNiaoTeam/controller"
	"CaiNiaoTeam/middleWare"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", middleWare.Authentication, controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.GET("/favorite/list/", middleWare.Authentication, controller.FavoriteList)
	apiRouter.POST("/publish/action/", controller.Publish)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleWare.Authentication, controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
