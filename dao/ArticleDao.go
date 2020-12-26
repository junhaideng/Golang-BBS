package dao

import "bbs/model"

func FindAllArticles() ([]*model.Article, error) {
	var articles []*model.Article
	if err := DB.Model(&model.Article{}).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func FindArticleById(id string) (*model.Article, error) {
	var article = &model.Article{}
	if err := DB.Model(&model.Article{}).Where("id = ?", id).Find(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func FindArticleReplyByArticleId(id string) ([]*model.Reply, error) {
	var reply []*model.Reply
	if err := DB.Model(&model.Reply{}).Where("article_id = ?", id).Find(&reply).Error; err != nil {
		return nil, err
	}
	return reply, nil
}

type hotArticle struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Read  uint   `json:"read"`
}

func FindHotArticleWithLimitAndOffset(page int, perPage int) ([]*hotArticle, error) {
	var articles []*hotArticle
	if err := DB.Model(&model.Article{}).Offset(page - 1).Limit(perPage).Select("id, title, `read`").Order("`read` desc").Scan(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func FindArticlesByUsername(username string) ([]*model.Article, error) {
	var articles []*model.Article
	if err := DB.Model(&model.Article{}).Where("username = ?", username).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func CreateArticle(article model.Article) error {
	tx := DB.Begin()
	if err := tx.Model(&model.Article{}).Create(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func FindArticleByString(query string) ([]*article, error) {
	var articles []*article
	if err := DB.Model(&model.Article{}).Where("title LIKE ?", "%"+query+"%").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// 删除文章
func DeleteArticle(username string, id string) error {
	tx := DB.Begin()
	if err := tx.Where("username = ? AND id = ?", username, id).Delete(&model.Article{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func AddArticleCommentsNum(articleId string, num uint) error {
	article, err := FindArticleById(articleId)
	if err != nil {
		return err
	}
	tx := DB.Begin()
	article.Comments += num
	if err := tx.Save(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func AddArticleReadNum(articleId string, num uint) error {
	article, err := FindArticleById(articleId)
	if err != nil {
		return err
	}
	tx := DB.Begin()
	article.Read += num
	if err := tx.Save(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
