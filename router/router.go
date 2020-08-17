package router

import (
	"net/http"

	"github.com/fengjunhua/devops-gin/handler"
	"github.com/gin-gonic/gin"
)

//SetupRouter is a function
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	router.POST("/POST", func(context *gin.Context) {

		context.String(http.StatusOK, "hello gin post method")

	})

	router.GET("/index", handler.Index)

	return router
}
