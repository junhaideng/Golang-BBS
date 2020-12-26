package controller

import (
	"bbs/dao"
	"bbs/model"
	"bbs/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Login(c *gin.Context) {
	service.Login(c)
}

func Register(c *gin.Context) {
	service.Register(c)
}

func GetMessage(c *gin.Context) {
	// 其实没有必要这样检查，因为中间件已经进行验证了
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": ErrorUnauthorized,
		})
		return
	}
	messages, err := dao.GetMessageByUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})
		return
	}
	c.JSON(http.StatusOK, messages)
}

// 标记信息已读
func ReadMessage(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}

	username, _ := c.Get("username")
	if err := dao.SetMessageRead(username.(string), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "标记成功",
	})
}

// 获取到未读信息数目
func GetUnreadMessageNum(c *gin.Context) {
	username, _ := c.Get("username")
	num := dao.GetUnreadMessageNumByUsername(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"data": num,
	})
}

func DeleteMessage(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	username, _ := c.Get("username")
	if err := dao.DeleteMessage(username.(string), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func GetUserInfo(c *gin.Context) {
	var user, exist = c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    user,
	})
}

// 获取用户登录信息
func GetLoginLog(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	username := user.(model.User).Username
	log, err := dao.FindLoginLogByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 该错误应该是由于username不存在导致的
			"message": "暂无对应信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
		"data":    log,
	})
}

func UpdateUserInfo(c *gin.Context) {
	var userInfo dao.UserInfo
	if err := c.Bind(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数有误",
		})
		return
	}
	fmt.Printf("%#v", userInfo)

	u, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	if err := dao.UpdateUser(u.(model.User), userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户信息更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户信息更新成功",
		})
	}

}

func UpdateAvatar(c *gin.Context) {
	service.UploadAvatar(c)
}

func GetAvatar(c *gin.Context) {
	username, exist := c.GetQuery("username")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数不正确",
		})
		return
	}
	path, err := dao.FindAvatarByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "无对应用户",
		})
		return
	}
	c.File(path.Path)
}

type PostReq struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Type    string `form:"type"`
}

// 发帖
func Post(c *gin.Context) {
	var req PostReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数不正确",
		})
		return
	}
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请先进行登录",
		})
		return
	}
	var article = model.Article{
		Username: username.(string),
		Title:    req.Title,
		Content:  req.Content,
		Type:     req.Type,
	}
	err := dao.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

// 修改密码
func ChangePassword(c *gin.Context) {
	service.ChangePassword(c)
}

func DeleteAccount(c *gin.Context) {
	service.DeleteAccount(c)
}

func GetMessageSettings(c *gin.Context) {
	username := c.GetString("username")
	setting, err := dao.GetMessageSettingsByUsername(username)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"username": username,
			"msg":      err.Error(),
		}).Error("get message settings error")

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})
		return
	}
	c.JSON(http.StatusOK, setting)
}

// 修改信息设置
func ChangeMessageSettings(c *gin.Context) {
	type param struct {
		Comment bool `form:"comment"`
		Like    bool `form:"like"`
		Star    bool `form:"star"`
	}
	req := param{}
	fmt.Println(req)

	if err := c.ShouldBind(&req); err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": err,
		}).Error("bind parameter error")

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	settings := map[string]interface{}{
		"username": c.GetString("username"),
		"comment":  req.Comment,
		"star":     req.Star,
		"like":     req.Like,
	}
	if err := dao.UpdateMessageSettingsByUsername(settings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})

		logrus.WithFields(logrus.Fields{
			"request":  req,
			"settings": settings,
			"msg":      err,
		}).Error("update message settings error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}
