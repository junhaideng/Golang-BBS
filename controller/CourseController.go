package controller

import (
	"bbs/dao"
	"bbs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllSchool(c *gin.Context) {
	schools, err := dao.GetAllSchool()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": schools,
	})
}

func GetCourse(c *gin.Context) {
	typ := c.PostForm("type")
	school := c.PostForm("school")

	courses, err := dao.GetCourseBySchoolAndType(school, typ)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})

}

func GetCourseDetail(c *gin.Context) {
	id := c.PostForm("id")
	school, err := dao.FindCourseById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": school,
	})
}

func GetCourseComments(c *gin.Context) {
	id := c.PostForm("id")
	comment, err := dao.FindCourseCommentByCourseId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func WriteCourseComment(c *gin.Context) {
	courseId, _ := strconv.ParseUint(c.PostForm("course_id"), 10, 64)
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "权限不足",
		})
		return
	}
	var comment = &model.CourseComment{
		Comment:  c.PostForm("comment"),
		CourseId: uint(courseId),
		Username: user.(model.User).Username,
	}

	if err := dao.CreateCourseComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "评论异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "评论成功",
	})
}

func GetCourseByValue(c *gin.Context) {
	value := c.PostForm("value")
	courses, err := dao.GetCourseByValue(value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取数据失败: ",
		})
	}
	c.JSON(http.StatusOK, courses)
}
