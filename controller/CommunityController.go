package controller

import (
	"bbs/dao"
	"bbs/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllArticles(c *gin.Context) {
	files, err := dao.FindAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("获取文章失败: %v", err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": files,
		})
	}
}

func GetHotArticle(c *gin.Context) {
	page, err := strconv.Atoi(c.PostForm("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	articles, err := dao.FindHotArticleWithLimitAndOffset(page, HotArticlePerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": articles,
	})
}

func GetArticleDetail(c *gin.Context) {
	id := c.PostForm("id")

	article, err := dao.FindArticleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": article,
		})
	}
}

func GetArticleReply(c *gin.Context) {
	id := c.PostForm("id")
	replies, err := dao.FindArticleReplyByArticleId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": replies,
		})
	}
}

func GetJwcNotice(c *gin.Context) {
	service.GetJwcNotice(c)
}
