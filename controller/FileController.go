package controller

import (
	"bbs/dao"
	"bbs/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func GetFilesByUser(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	files, err := dao.FindFileByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "没有查找到对应的文件",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
		"data":    files,
	})
}

// 删除文件
func DeleteUserFile(c *gin.Context) {
	type param struct {
		Ids []int `form:"ids" json:"ids"`
	}
	var req = param{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
	}
	username := c.GetString("username")
	for _, id := range req.Ids {
		if err := dao.DeleteFileById(username, id); err != nil {
			logrus.WithFields(logrus.Fields{
				"id":  id,
				"msg": err,
			}).Error("delete file error")
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func DownloadFile(c *gin.Context) {
	id, exist := c.GetQuery("file_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	file, err := dao.FindFileByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.File(file.Path)
	if err := dao.AddFileDownloadTimes(id, 1); err != nil {
		fmt.Println(err)
	}
}
