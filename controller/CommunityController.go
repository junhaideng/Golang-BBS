package controller

import (
	"bbs/dao"
	"bbs/model"
	"bbs/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllArticles(c *gin.Context) {
	articles, err := dao.FindAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("获取文章失败: %v", err),
		})
	} else {
		c.JSON(http.StatusOK, articles)
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
	c.JSON(http.StatusOK, articles)
}

func GetArticleDetail(c *gin.Context) {
	id := c.PostForm("id")

	article, err := dao.FindArticleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, article)
	if err := dao.AddArticleReadNum(id, 1); err != nil {
		fmt.Println(err)
	}
}

func GetArticleReply(c *gin.Context) {
	type response struct {
		*model.Reply
		Comment []*model.Comment `json:"comment"`
	}
	var res []response
	id := c.PostForm("id")
	replies, err := dao.FindArticleReplyByArticleId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
	}
	for _, reply := range replies {
		replyId := reply.Id
		comments, err := dao.FindCommentByReplyId(fmt.Sprintf("%d", replyId))
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, response{
			Reply:   reply,
			Comment: comments,
		})
	}

	c.JSON(http.StatusOK, res)

}

func GetJwcNotice(c *gin.Context) {
	service.GetJwcNotice(c)
}

// 在文章下面进行回复
func ReplyArticle(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请先进行登录",
		})
		return
	}
	type res struct {
		ID    uint   `form:"id" binding:"required"`
		Reply string `form:"reply"`
	}
	param := &res{}
	if err := c.ShouldBind(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	reply := model.Reply{
		ArticleId: param.ID,
		Reply:     param.Reply,
		Username:  username.(string),
	}
	if err := dao.CreateReply(reply); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建回复失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "回复成功",
	})
}

// 给对应的评论下面进行回复
func CommentReply(c *gin.Context) {
	type req struct {
		Comment string `form:"comment" binding:"required"`
		ReplyID string `form:"reply_id" binding:"required"`
		URL     string `form:"url" binding:"required"` // 给文章作者进行回复时使用 TODO
	}
	var param req
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": ErrorBadRequest,
		})
		return
	}
	username, _ := c.Get("username")
	if err := dao.CreateCommentToReply(username.(string), param.ReplyID, param.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": ErrorInternalServer,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "评论成功",
	})
}
