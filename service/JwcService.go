package service

import (
	"bbs/dao"
	"bbs/model"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
)

type Notice struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item" json:"item"`
	Title   string   `xml:"title" json:"title"`
	Link    string   `xml:"link" json:"link"`
	PubDate string   `xml:"pubDate" json:"pubDate"`
}

var db = dao.DB

func GetJwcNotice(c *gin.Context) {
	resp, err := http.Get("http://jwc.sjtu.edu.cn/rss/rss_notice.aspx?SubjectID=198015&TemplateID=221027")

	if err != nil {
		fmt.Println("http.Get err=", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取数据失败",
		})
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err=", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "读取数据失败",
		})
		return
	}
	var notice Notice
	decoder := xml.NewDecoder(bytes.NewReader(buf))
	decoder.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		return charset.NewReaderLabel(c, i)
	}

	err = decoder.Decode(&notice)

	if err != nil {
		fmt.Println("反序列化错误: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "数据处理异常",
		})
		return
	}

	//for _, item := range notice.Channel.Item[:10] {
	//	insertJwcNotice(item)
	//}
	c.JSON(http.StatusOK, notice.Channel.Item[:10])
}

// 将查询到的信息插入到数据库
// 由于每次都需要请求的都不一样
// 所以其实没有这个必要
func insertJwcNotice(item Item) {
	tx := db.Begin()
	notice := model.JwcNotice{
		Title:   item.Title,
		Link:    item.Link,
		PubDate: item.PubDate,
	}
	if err := tx.Model(&model.JwcNotice{}).Create(&notice).Error; err != nil {
		tx.Rollback()
		fmt.Println("数据插入失败.")
		return
	}
	tx.Commit()
}
