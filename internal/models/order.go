package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	CreatorBase
	Status                uint      `json:"status" comment:"订单状态 0-默认 10-已预定 20-待支付 30-已支付，待发货 40-待收货 50-待评价 60-已评价，订单已完成 70-待退款，审核中 80-退款成功 90-退款失败"`
	EstimatedAmount       float64   `json:"estimatedAmount" comment:"预估金额"`
	RealAmount            float64   `json:"realAmount" comment:"实际金额"`
	RefundAmount          float64   `json:"refundAmount" comment:"应退金额"`
	RefundAcceptTime      time.Time `json:"refundAcceptTime" comment:"退款申请时间"`
	RefundAcceptIntroduce string    `json:"RefundAcceptIntroduce" comment:"退款申请说明"`
	PayWay                uint      `json:"payWay" comment:"支付方式 0-支付宝 1-微信 2-网银"`
	Source                uint      `json:"source" comment:"订单来源 0-app 1-web 2-小程序"`
	Evaluation            string    `json:"evaluation" comment:"用户评价"`
}

func (t *Order) TableName() string {
	return "orders"
}

func (t *Order) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *Order) Updates(db *gorm.DB, value map[string]interface{}, conditions interface{}, arg ...interface{}) error {
	return db.Model(&t).Where(conditions, arg...).Updates(value).Error
}

func (t *Order) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (*Order, error) {
	var order = new(Order)
	err := db.Model(&t).Where(query, args...).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (t *Order) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (orders []*Order, err error) {
	err = db.Model(&t).Where(query, args...).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
