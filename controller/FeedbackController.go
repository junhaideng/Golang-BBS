package controller

import (
	"bbs/dao"
	"bbs/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SubmitFeedback(c *gin.Context) {
	type req struct {
		Title   string `form:"title"`
		Email   string `form:"email"`
		Content string `form:"content"`
	}
	param := req{}
	if err := c.ShouldBind(&param); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": err.Error(),
		}).Error("bind param error")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	username, _ := c.Get("username")
	feedback := model.Feedback{
		Content:  param.Content,
		Email:    param.Email,
		Title:    param.Title,
		Username: username.(string),
		Active:   false,
	}
	if err := dao.CreateFeedback(feedback); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg":      err.Error(),
			"feedback": feedback,
		}).Error("insert feedback error")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建反馈失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "感谢反馈",
	})
}
