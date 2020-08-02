package controller

import (
	"bbs/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	service.Login(c)
}

func Register(c *gin.Context) {
	service.Register(c)
}
func GetUnreadMessageNum(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": 2,
	})
}
