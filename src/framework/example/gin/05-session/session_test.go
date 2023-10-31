package _5_session

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestSession(t *testing.T) {

	r := gin.Default()

	// 使用 Cookie 作为存储介质，传入加密密钥
	store := cookie.NewStore([]byte("secret"))
	// 使用 sessions 中间件
	r.Use(sessions.Sessions("mysession", store))

	// 设置 Session
	r.GET("/set-session", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("username", "JohnDoe")
		// 设置session过期时间
		session.Options(sessions.Options{MaxAge: 10})
		session.Save()
		c.String(http.StatusOK, "Session has been set")
	})

	// 获取 Session
	r.GET("/get-session", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.String(http.StatusOK, "Session not found")
			return
		}
		c.String(http.StatusOK, "Username: %s", username)
	})

	r.Run(":8080")
}

func TestRedisSessionStore(t *testing.T) {

	r := gin.Default()

	// 使用 Cookie 作为存储介质，传入加密密钥
	//store := cookie.NewStore([]byte("secret"))
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		fmt.Println("redis connect err", err)
	}
	// 使用 sessions 中间件
	r.Use(sessions.Sessions("mysession", store))

	// 设置 Session
	r.GET("/set-session", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("username", "JohnDoe")
		// 设置session过期时间
		session.Options(sessions.Options{MaxAge: 10})
		session.Save()
		c.String(http.StatusOK, "Session has been set")
	})

	// 获取 Session
	r.GET("/get-session", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.String(http.StatusOK, "Session not found")
			return
		}
		c.String(http.StatusOK, "Username: %s", username)
	})

	r.Run(":8080")
}
