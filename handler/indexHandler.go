package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//Index 返回一个index界面的函数
func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "嘿嘿嘿" + strings.ToLower(context.Request.Method) + " method",
	})
}
