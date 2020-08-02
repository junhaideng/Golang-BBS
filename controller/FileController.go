package controller

import (
	"bbs/dao"
	"bbs/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	service.UploadFile(c)
}

func GetAllFiles(c *gin.Context) {
	files, err := dao.FindAllFiles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "查询失败",
		})
	}
	c.JSON(http.StatusOK, files)
}
