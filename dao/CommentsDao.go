package dao

import (
	"bbs/model"
	"strconv"
)

// 通过replyId查找对应的回复
func FindCommentByReplyId(replyId string) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := DB.Model(&model.Comment{}).Where("reply_id = ? ", replyId).Scan(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// 创建对应回复Reply的评论Comment
func CreateCommentToReply(username, replyId, comment string) error {
	if err := AddReplyCommentNum(replyId, 1); err != nil {
		return err
	}
	tx := DB.Begin()
	id, err := strconv.ParseUint(replyId, 10, 32)
	if err != nil {
		return err
	}
	comm := &model.Comment{
		Comment:  comment,
		ReplyId:  uint(id),
		Username: username,
	}
	if err := tx.Create(comm).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
