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
	var article *model.Article
	if err := DB.Model(&model.Article{}).Where("id = ?", id).Find(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func FindArticleReplyByArticleId(id string) ([]*model.Reply, error) {
	var reply []*model.Reply
	if err := DB.Model(&model.Reply{}).Where("id = ?", id).Find(&reply).Error; err != nil {
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
