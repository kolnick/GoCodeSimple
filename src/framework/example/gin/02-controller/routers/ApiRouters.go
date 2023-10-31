package routers

import (
	"GoCodeSimple/src/framework/example/gin/02-controller/controllers"
	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", controllers.UserIndex)
		apiGroup.GET("/plist", controllers.UserList)
	}
}
