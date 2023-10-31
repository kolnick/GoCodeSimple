package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexRoutersInit(r *gin.Engine) {
	group := r.Group("/")
	{
		group.GET("/", func(g *gin.Context) {
			g.HTML(http.StatusOK, "defaults/index.html", gin.H{
				"msg": "我是一个msg",
			})
		})
		group.GET("/news", func(g *gin.Context) {
			g.String(http.StatusOK, "新闻")
		})
		group.GET("/index", func(context *gin.Context) {
			context.String(http.StatusOK, "首页")
		})
	}
}
