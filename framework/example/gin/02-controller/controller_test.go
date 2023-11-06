package _2_controller

import (
	routers2 "GoCodeSimple/framework/example/gin/02-controller/routers"
	"github.com/gin-gonic/gin"
	"testing"
)

var r *gin.Engine

func init() {
	// 创建一个默认的路由引擎
	r = gin.Default()
}

func TestGroupRouter(t *testing.T) {
	routers2.IndexRoutersInit(r)
	routers2.ApiRouterInit(r)
	r.Run()
}
