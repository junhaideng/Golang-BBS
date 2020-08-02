package service

import (
	"bbs/dao"
	"bbs/model"
	"bbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := dao.FindUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
		})
		return
	}

	token, err := utils.ReleaseToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
		},
		"message": "登录成功",
	})
}

type RegisterUser struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Gender   string `form:"gender"`
	Academy  string `form:"academy"`
	Grade    string `form:"grade"`
}

func Register(c *gin.Context) {
	var registerUser RegisterUser
	if err := c.ShouldBind(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数不正确",
		})
		return
	}
	fmt.Println(registerUser)
	user, err := dao.FindUserByUsername(registerUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
	}
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名已存在，请更换用户名重新注册",
		})
		return
	}

	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统错误",
		})
		return
	}

	user = &model.User{
		Username: registerUser.Username,
		Password: string(encryptPwd),
		Academy:  registerUser.Academy,
		Email:    registerUser.Email,
		Grade:    registerUser.Grade,
		Gender:   registerUser.Gender,
	}

	if err := dao.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户创建失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户创建成功",
		})
	}

}
