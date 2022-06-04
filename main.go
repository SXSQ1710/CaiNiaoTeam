package main

import (
	"CaiNiaoTeam/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitCreatTable()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//10000
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQyMzc2NzQsInVzZXJfaWQiOiIxMDAwMCJ9.F-6If3xBBSnNrE04u8rBASY9MHBhjspr1-5qYURKiMc
}
