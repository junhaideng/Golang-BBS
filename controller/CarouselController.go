package controller

import (
	"bbs/dao"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func GetCarousel(c *gin.Context) {
	// 首先获取需要的个数
	num := c.DefaultPostForm("num", "3")
	// 保存数据
	// 向数据库中查找对应的数据
	carousel, err := dao.GetCarouselIdByNum(num)
	// 返回json类型数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": carousel,
	})
}
