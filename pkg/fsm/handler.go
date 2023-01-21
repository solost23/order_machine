package fsm

import (
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/order_machine"
)

var (
	// 创建订单(此业务应该写到service的订单创建接口中)
	handlerCreate = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		return order_machine.OrderStatus_StatusReserved, nil
	})
	// 确认订单
	handlerConfirm = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]确认了订单[%d] \n", opt.Order.CreatorId, opt.Order.ID)
		return order_machine.OrderStatus_StatusWaitPayment, nil
	})
	// 修改订单 - 传入订单信息修改
	handlerModify = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]修改了订单信息 \n", opt.Order.CreatorId)
		return order_machine.OrderStatus_StatusReserved, nil
	})
	// 支付订单
	handlerPay = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]支付订单[%d]成功，支付金额[%v] \n", opt.Order.CreatorId, opt.Order.ID, opt.Order.RealAmount)
		return order_machine.OrderStatus_StatusAlreadyPaymentWaitSendGoods, nil
	})
	// 订单发送
	handlerSend = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]的订单[%d]已发送 \n", opt.Order.CreatorId, opt.Order.ID)
		return order_machine.OrderStatus_StatusWaitAcceptGoods, nil
	})
	// 用户收货
	handlerAccept = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]已收到订单[%d] \n", opt.Order.CreatorId, opt.Order.ID)
		return order_machine.OrderStatus_StatusWaitEvaluation, nil
	})
	// 用户评价
	handlerEvaluation = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]评价了订单[%d], 评价信息为[%s] \n", opt.Order.CreatorId, opt.Order.ID, opt.Order.Evaluation)
		return order_machine.OrderStatus_StatusAlreadyEvaluationOrderOver, nil
	})
	// 用户申请退款
	handlerRefundAccept = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]对订单[%d]申请了退款, 退款金额为[%v] \n", opt.Order.CreatorId, opt.Order.ID, opt.Order.RealAmount)
		return order_machine.OrderStatus_StatusWaitRefundChecking, nil
	})
	// 用户取消退款 规定回到已支付，待发货状态
	handlerRefundCancel = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("用户[%d]对订单[%d]取消了退款 \n", opt.Order.CreatorId, opt.Order.ID)
		return order_machine.OrderStatus_StatusAlreadyPaymentWaitSendGoods, nil
	})
	// 后台同意退款
	handlerRefundConfirm = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("后台同意对用户[%d]的订单[%d]退款，退款金额为[%v] \n", opt.Order.CreatorId, opt.Order.ID, opt.Order.RealAmount)
		return order_machine.OrderStatus_StatusRefundSuccess, nil
	})
	// 后台拒绝退款
	handlerRefundReject = Handler(func(opt *Opt) (order_machine.OrderStatus, error) {
		fmt.Printf("后台拒绝对用户[%d]的订单[%d]退款 \n", opt.Order.CreatorId, opt.Order.ID)
		return order_machine.OrderStatus_StatusRefundFailed, nil
	})
)
