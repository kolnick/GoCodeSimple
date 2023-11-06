package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndex(c *gin.Context) {
	c.String(http.StatusOK, "用户首页")
}
func UserList(c *gin.Context) {
	c.String(http.StatusOK, "用户列表")
}
