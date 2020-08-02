package dao

import (
	"bbs/model"
)

type schoolInfo struct {
	School string `json:"school"`
}

func GetAllSchool() ([]*schoolInfo, error) {
	var schools []*schoolInfo
	if err := DB.Model(model.Course{}).Select("distinct(school)").Scan(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}

type courseName struct {
	CourseName string `json:"course_name"`
}

func GetCourseByValue(value string) ([]*courseName, error) {
	var courses []*courseName
	if err := DB.Model(model.Course{}).Select("distinct(course_name)").Where("course_name like ?", "%"+value+"%").Scan(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

type courseInfo struct {
	Id         uint   `json:"id"`
	CourseName string `json:"course_name"`
	Teacher    string `json:"teacher"`
	School     string `json:"school"`
}

func GetCourseBySchoolAndType(school string, typ string) ([]*courseInfo, error) {
	var courses []*courseInfo
	if err := DB.Model(&model.Course{}).Select("id, school, course_name, teacher").Where("school like ? AND type like ?", "%"+school, "%"+typ).Scan(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func FindCourseById(id string) (*model.Course, error) {
	var course *model.Course
	if err := DB.Model(&model.Course{}).First(&course, id).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func CreateCourseComment(comment *model.CourseComment) error {
	tx := DB.Begin()
	if err := tx.Model(&model.CourseComment{}).Create(comment).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func FindCourseCommentByCourseId(id string)([]*model.CourseComment, error){
	var comment []*model.CourseComment
	if err := DB.Model(&model.CourseComment{}).Where("course_id = ?", id).Find(&comment).Error; err != nil{
		return nil, err
	}
	return comment, nil
}
