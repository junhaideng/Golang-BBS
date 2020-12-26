package main

import (
	_ "bbs/config" // 引进来只是为了初始化
	"bbs/controller"
	_ "bbs/log"
	"bbs/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	v := r.Group("/api")

	v1 := v.Group("/carousel")
	{
		v1.POST("/get", controller.GetCarousel)
		v1.GET("/get", controller.ServeCarousel)
	}

	v2 := v.Group("/course")
	{
		v2.GET("/get_all_school", controller.GetAllSchool)
		v2.POST("/get_course", controller.GetCourse)
		v2.POST("/get", controller.GetCourseDetail)
		v2.POST("/get_course_comments", controller.GetCourseComments)
		v2.POST("/write_course_comments", middleware.AuthMiddleware(), controller.WriteCourseComment)
		v2.POST("/get_course_by_value", controller.GetCourseByValue)
	}

	v3 := v.Group("/community")
	{
		v3.GET("/get_all", controller.GetAllArticles)
		v3.POST("/hot", controller.GetHotArticle)
		v3.POST("/get_article", controller.GetArticleDetail)
		v3.POST("/get_reply", controller.GetArticleReply)
		v3.GET("/jwc", controller.GetJwcNotice)
		v3.Use(middleware.AuthMiddleware())
		v3.POST("/reply", controller.ReplyArticle)
		v3.POST("/reply/comment", controller.CommentReply)
	}

	v4 := v.Group("/user")
	{
		v4.POST("/login", controller.Login)
		v4.POST("/register", controller.Register)
		v4.GET("/avatar/get", controller.GetAvatar)
		v4.Use(middleware.AuthMiddleware())
		v4.POST("/message", controller.GetMessage)
		v4.POST("/message/read", controller.ReadMessage)
		v4.POST("/message/unread", controller.GetUnreadMessageNum)
		v4.POST("/message/delete", controller.DeleteMessage)
		v4.POST("/info", controller.GetUserInfo)
		v4.POST("/updateUserInfo", controller.UpdateUserInfo)
		v4.POST("/avatar/upload", controller.UpdateAvatar)
		v4.GET("/article", controller.GetArticleByUsername)
		v4.POST("/article/delete", controller.DeleteArticle)
		v4.GET("/loginlog", controller.GetLoginLog)
		v4.POST("/post", controller.Post)
		v4.GET("/files", controller.GetFilesByUser)
		v4.POST("/files/delete", controller.DeleteUserFile) // TODO
		v4.POST("/change_password", controller.ChangePassword)
		v4.POST("/delete", controller.DeleteAccount)

		//
		v4.GET("/messageSettings", controller.GetMessageSettings)
		v4.POST("/changesettings", controller.ChangeMessageSettings)
	}

	v5 := v.Group("/file")
	{
		v5.GET("/get_all", controller.GetAllFiles)
		v5.POST("/upload", middleware.AuthMiddleware(), controller.UploadFile)
		v5.GET("/download", controller.DownloadFile)
	}

	v6 := v.Group("/search")
	{
		v6.POST("", controller.SearchArticle)
	}

	v7 := v.Group("/feedback")
	{
		v7.Use(middleware.AuthMiddleware())
		v7.POST("/submit", controller.SubmitFeedback)
	}

	if err := r.Run(":8000"); err != nil {
		logrus.Error(err)
	}
}
