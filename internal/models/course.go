package models

import "gorm.io/gorm"

type Course struct {
	CreatorBase
	Name   string  `json:"name" comment:"课程名称"`
	Amount float64 `json:"amount" comment:"课程金额"`
	Stock  uint64  `json:"stock" comment:"课程存量"`
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

func (t *Course) WhereAll(db *gorm.DB, query string, args ...interface{}) (results []*Course, err error) {
	err = db.Model(&t).Where(query, args...).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (t *Course) Update(db *gorm.DB, values interface{}, conditions string, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(values).Error
}
