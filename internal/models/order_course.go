package models

import "gorm.io/gorm"

type OrderCourse struct {
	CreatorBase
	OrderId  uint `json:"orderId" comment:"订单 ID"`
	CourseId uint `json:"courseId" comment:"课程 ID"`
}

func (t *OrderCourse) TableName() string {
	return "order_courses"
}

func (t *OrderCourse) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *OrderCourse) Delete(db *gorm.DB, conditions string, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Delete(&t).Error
}

func (t *OrderCourse) WhereAll(db *gorm.DB, query string, args ...interface{}) (results []*OrderCourse, err error) {
	err = db.Model(&t).Where(query, args...).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (t *OrderCourse) WhereCount(db *gorm.DB, query string, args ...interface{}) (count int64, err error) {
	err = db.Model(&t).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
