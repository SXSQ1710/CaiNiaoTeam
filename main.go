package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")


	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", userController.Register)
}
