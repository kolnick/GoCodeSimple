package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

var r *gin.Engine

func init() {
	// 创建一个默认的路由引擎
	r = gin.Default()
}

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		//AllowAllOrigins: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", http.MethodDelete, http.MethodHead, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true,
	}
	return cors.New(cfg)
}

func TestCors(t *testing.T) {
	r.Use(Cors())
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/userList", func(g *gin.Context) {
			g.String(http.StatusOK, "用户列表")
		})
		apiGroup.GET("/plist", func(g *gin.Context) {
			g.String(http.StatusOK, "列表")
		})
	}
	r.Run()
}
