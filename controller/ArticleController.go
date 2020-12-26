package controller

import (
	"bbs/dao"
	"bbs/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticleByUsername(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
	}
	data, err := dao.FindArticlesByUsername(user.(model.User).Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    data,
	})
}

func DeleteArticle(c *gin.Context) {
	id, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	username, _ := c.Get("username")
	err := dao.DeleteArticle(username.(string), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func SearchArticle(c *gin.Context) {
	query := c.PostForm("q")
	if len(query) == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	data, err := dao.FindArticleByString(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
