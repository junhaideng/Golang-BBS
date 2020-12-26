package dao

import "bbs/model"

type carouselInfo struct {
	Id uint `json:"id"`
}

func GetCarouselIdByNum(num string) ([]*carouselInfo, error) {
	var carousel []*carouselInfo
	err := DB.Model(&model.Carousel{}).Limit(num).Order("time desc").Select("id").Where("active = ?", 1).Scan(&carousel).Error
	if err != nil {
		return nil, err
	}
	return carousel, nil
}

func GetCarouselById(id string) (*model.Carousel, error) {
	var carousel = &model.Carousel{}

	err := DB.Model(&model.Carousel{}).Where("id = ?", id).Scan(&carousel).Error
	if err != nil {
		return nil, err
	}
	return carousel, nil
}
