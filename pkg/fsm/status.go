package fsm

import "github.com/solost23/protopb/gen/go/protos/order_machine"

// statusText 定义订单状态文案
var statusText = map[order_machine.OrderStatus]string{
	order_machine.OrderStatus_StatusDefault:                     "默认",
	order_machine.OrderStatus_StatusReserved:                    "已预定",
	order_machine.OrderStatus_StatusWaitPayment:                 "待支付",
	order_machine.OrderStatus_StatusAlreadyPaymentWaitSendGoods: "已支付，待发货",
	order_machine.OrderStatus_StatusWaitAcceptGoods:             "待收货",
	order_machine.OrderStatus_StatusWaitEvaluation:              "待评价",
	order_machine.OrderStatus_StatusAlreadyEvaluationOrderOver:  "已评价，订单已完成",
	order_machine.OrderStatus_StatusWaitRefundChecking:          "待退款，审核中",
	order_machine.OrderStatus_StatusRefundSuccess:               "退款成功",
	order_machine.OrderStatus_StatusRefundFailed:                "退款失败",
}

// stateEvent 定义订单状态对应的可操作事件
var StatusEvent = map[order_machine.OrderStatus][]order_machine.OrderEvent{
	order_machine.OrderStatus_StatusDefault:                     {order_machine.OrderEvent_EventCreate},
	order_machine.OrderStatus_StatusReserved:                    {order_machine.OrderEvent_EventConfirm},
	order_machine.OrderStatus_StatusWaitPayment:                 {order_machine.OrderEvent_EventModify, order_machine.OrderEvent_EventPay},
	order_machine.OrderStatus_StatusAlreadyPaymentWaitSendGoods: {order_machine.OrderEvent_EventSend, order_machine.OrderEvent_EventRefundAccept},
	order_machine.OrderStatus_StatusWaitAcceptGoods:             {order_machine.OrderEvent_EventAccept, order_machine.OrderEvent_EventRefundAccept},
	order_machine.OrderStatus_StatusWaitEvaluation:              {order_machine.OrderEvent_EventEvaluation, order_machine.OrderEvent_EventRefundAccept},
	//StatusAlreadyEvaluationOrderOver:  {},
	order_machine.OrderStatus_StatusWaitRefundChecking: {order_machine.OrderEvent_EventRefundConfirm, order_machine.OrderEvent_EventRefundCancel, order_machine.OrderEvent_EventRefundReject},
	//StatusRefundSuccess:               {},
	//StatusRefundFailed:                {},
}

func StatusText(status order_machine.OrderStatus) string {
	return statusText[status]
}
