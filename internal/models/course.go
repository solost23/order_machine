package models

import "gorm.io/gorm"

type Course struct {
	CreatorBase
	Name   string  `json:"name" comment:"课程名称"`
	Amount float64 `json:"amount" comment:"课程金额"`
}

func (t *Course) TableName() string {
	return "courses"
}

func (t *Course) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *Course) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (*Course, error) {
	var course = new(Course)
	err := db.Model(&t).Where(query, args...).First(&course).Error
	if err != nil {
		return nil, err
	}
	return course, nil
}
