package fsm

import (
	"github.com/solost23/go_interface/gen_go/order_machine"
)

var (
	// 创建订单(此业务应该写到service的订单创建接口中)
	handlerCreate = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusReserved, nil
	})
	// 确认订单
	handlerConfirm = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusWaitPayment, nil
	})
	// 修改订单 - 传入订单信息修改
	handlerModify = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusReserved, nil
	})
	// 支付订单
	handlerPay = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusAlreadyPaymentWaitSendGoods, nil
	})
	// 订单发送
	handlerSend = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusWaitAcceptGoods, nil
	})
	// 用户收货
	handlerAccept = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusWaitEvaluation, nil
	})
	// 用户评价
	handlerEvaluation = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusAlreadyEvaluationOrderOver, nil
	})
	// 用户申请退款
	handlerRefundAccept = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusWaitRefundChecking, nil
	})
	// 用户取消退款 规定回到已支付，待发货状态
	handlerRefundCancel = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusAlreadyPaymentWaitSendGoods, nil
	})
	// 后台同意退款
	handlerRefundConfirm = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusRefundSuccess, nil
	})
	// 后台拒绝退款
	handlerRefundReject = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusRefundFailed, nil
	})
)
