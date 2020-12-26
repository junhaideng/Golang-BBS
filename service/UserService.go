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
		if err == dao.ErrUserNotExist {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "密码错误",
		})
		return
	}

	token, err := utils.ReleaseToken(username)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "系统异常",
		})
		return
	}
	info, err := utils.GetLocationByIp(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "请求资源失败",
		})
		return
	}
	log := &model.LoginLog{
		Address:  info.Content.Address,
		Ip:       c.ClientIP(),
		Username: username,
	}
	if err := dao.CreateLoginLog(log); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "记录日志失败",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
		},
		"code":    200,
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
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "参数不正确",
		})
		return
	}
	fmt.Println(registerUser)
	user, err := dao.FindUserByUsername(registerUser.Username)
	if err != dao.ErrUserNotExist && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "系统异常",
		})
		fmt.Println(err)
		return
	}

	if user != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "用户名已存在，请更换用户名重新注册",
		})
		return
	}

	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
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

	fmt.Printf("%#v\n", user)

	if err := dao.CreateUser(user); err != nil {
		if err == dao.ErrUserAlreadyExist {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "用户已存在",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "用户创建失败",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "用户创建成功",
		})
	}
}

type PasswordInfo struct {
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}

func ChangePassword(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	hashedPassword, err := dao.GetPasswordByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "修改失败",
		})
		return
	}
	var password PasswordInfo
	if err := c.ShouldBind(&password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数不正确",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password.OldPassword)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "密码不正确",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统异常",
		})
		return
	}

	if err := dao.UpdatePasswordByUsername(username.(string), string(hash)); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "密码修改错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

func DeleteAccount(c *gin.Context) {
	pwd := c.PostForm("password")

	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	password, err := dao.GetPasswordByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "异常",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码不正确",
		})
		return
	}
	err = dao.DeleteAccountByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "删除错误",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	}
}
