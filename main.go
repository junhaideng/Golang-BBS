package main

import (
	"bbs/common"
	"bbs/controller"
	"bbs/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(common.AvatarPath)
	r := gin.Default()
	v := r.Group("/api")

	v1 := v.Group("/carousel")
	{
		v1.POST("/get", controller.GetCarousel)
	}

	v2 := v.Group("/course")
	{
		v2.GET("/get_all_school", controller.GetAllSchool)
		v2.POST("/get_course", controller.GetCourse)
		v2.POST("/get", controller.GetCourseDetail)
		v2.POST("/get_course_comments", controller.GetCourseComments)
		v2.POST("/write_course_comments",middleware.AuthMiddleware(), controller.WriteCourseComment)
		v2.POST("/get_course_by_value", controller.GetCourseByValue)
	}

	v3 := v.Group("/community")
	{
		v3.GET("/get_all", controller.GetAllArticles)
		v3.POST("/hot", controller.GetHotArticle)
		v3.POST("/get_article", controller.GetArticleDetail)
		v3.POST("/get_reply", controller.GetArticleReply)
		v3.GET("/jwc", controller.GetJwcNotice)
	}

	v4 := v.Group("/user")
	{
		v4.POST("/login", controller.Login)
		v4.POST("/register", controller.Register)
		v4.POST("/message/unread", controller.GetUnreadMessageNum)

	}

	v5 := v.Group("/file")
	{
		v5.GET("/get_all", controller.GetAllFiles)
		v5.POST("/upload", middleware.AuthMiddleware(), controller.UploadFile)
	}

	r.Run(":8000")
}
