package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

var r *gin.Engine

func init() {
	// 创建一个默认的路由引擎
	r = gin.Default()
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

func TestGet(t *testing.T) {
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", sayHello)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
	// r.Run(":9090")
}

func TestStaticWeb(t *testing.T) {
	// 加载模板,放在配置路由前面
	//r.LoadHTMLGlob("templates/**/*")
	// 配置静态web目录，第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")
	r.GET("/hello", sayHello)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
	// r.Run(":9090")

}
