package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRouterInit(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/userList", func(g *gin.Context) {
			g.String(http.StatusOK, "用户列表")
		})
		apiGroup.GET("/plist", func(g *gin.Context) {
			g.String(http.StatusOK, "列表")
		})
	}
}
