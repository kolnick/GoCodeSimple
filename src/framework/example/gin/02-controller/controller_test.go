package _2_controller

import (
	"GoCodeSimple/src/framework/example/gin/02-controller/routers"
	"github.com/gin-gonic/gin"
	"testing"
)

var r *gin.Engine

func init() {
	// 创建一个默认的路由引擎
	r = gin.Default()
}

func TestGroupRouter(t *testing.T) {
	routers.IndexRoutersInit(r)
	routers.ApiRouterInit(r)
	r.Run()
}
