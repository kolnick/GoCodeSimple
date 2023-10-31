package _1_router

import (
	routers2 "GoCodeSimple/src/framework/example/gin/01-router/routers"
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
