package service

import (
	"bbs/common"
	"bbs/dao"
	"bbs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件上传异常",
		})
		return
	}
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	cwd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件保存异常",
		})
		return
	}

	savePath := filepath.Join(cwd, common.FileUploadPath)
	// 文件保存的位置
	savePath = filepath.Join(savePath, t+"_"+file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件保存异常: ",
		})
		return
	}
	desc := c.PostForm("description")
	typ := c.PostForm("type")

	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
		})
		return
	}

	var f = &model.File{
		Description: desc,
		Username:    user.(model.User).Username,
		Filename:    file.Filename,
		Path:        savePath,
		Type:        typ,
	}

	err = dao.CreateFile(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文件上传成功",
	})
}
