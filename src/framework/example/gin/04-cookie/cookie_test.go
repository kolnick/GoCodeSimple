package _4_cookie

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestCookie(t *testing.T) {

	r := gin.Default()

	r.GET("/set-cookie", func(c *gin.Context) {

		c.SetCookie("username", "JohnDoe", 3500, "/", "localhost", false, true)

		c.String(http.StatusOK, "Cookie has been set")
	})
	r.GET("/get-cookie", func(c *gin.Context) {

		username, err := c.Cookie("username")
		if err != nil {
			c.String(http.StatusOK, "Cookie not found")
			return
		}

		c.String(http.StatusOK, "Username:%s", username)
	})
	r.Run()
}
