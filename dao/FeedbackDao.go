package dao

import "bbs/model"

func CreateFeedback(feedback model.Feedback) error {
	tx := DB.Begin()
	if err := tx.Create(&feedback).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
