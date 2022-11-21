package fsm

import "order_machine/internal/models"

type Option func(*Opt)

// SendSMS 发送短信
type SendSMS func(mobile, content string) error

// Opt 定义 Handler 所需参数
type Opt struct {
	Order     *models.Order
	CourseNum uint

	HandlerSendSMS SendSMS
}

// 设置订单信息
func WithOrderInfo(order *models.Order) Option {
	return func(opt *Opt) {
		opt.Order = order
	}
}

// 设置商品数量
func WithOrderCourseNum(courseNum uint) Option {
	return func(opt *Opt) {
		opt.CourseNum = courseNum
	}
}

// WithHandlerSendSMS 设置发送短信
func WithHandlerSendSMS(sendSms SendSMS) Option {
	return func(opt *Opt) {
		opt.HandlerSendSMS = sendSms
	}
}
