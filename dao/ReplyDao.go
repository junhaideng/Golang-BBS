package dao

import (
	"bbs/model"
	"fmt"
)

func CreateReply(reply model.Reply) error {
	tx := DB.Begin()
	if err := tx.Model(&model.Article{}).Create(&reply).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := AddArticleCommentsNum(fmt.Sprintf("%d", reply.ArticleId), 1); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func FindReplyByID(replyId string) (*model.Reply, error) {
	var reply = &model.Reply{}
	if err := DB.Model(&model.Reply{}).Where("id = ?", replyId).Scan(reply).Error; err != nil {
		return nil, err
	}
	return reply, nil
}

func AddReplyCommentNum(replyId string, num uint) error {
	reply, err := FindReplyByID(replyId)
	if err != nil {
		return err
	}
	tx := DB.Begin()
	reply.Comments += num
	if err := tx.Save(&reply).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
