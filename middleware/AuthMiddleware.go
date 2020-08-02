package middleware

import (
	"bbs/dao"
	"bbs/model"
	"bbs/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}

		// 获取真正的token字符串
		tokenString = tokenString[7:]

		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}

		username := claims.Username
		db := dao.DB
		var user model.User
		db.Model(&model.User{}).Where("username = ?", username).First(&user)

		// 如果没有找到
		if user.Id == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足",
			})
			c.Abort()
			return
		}

		// 如果存在
		c.Set("user", user)
		c.Next()

	}
}
