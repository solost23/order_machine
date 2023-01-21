package fsm

import "github.com/solost23/protopb/gen/go/protos/order_machine"

// 定义事件对应的处理办法
var eventHandler = map[order_machine.OrderEvent]Handler{
	order_machine.OrderEvent_EventCreate:        handlerCreate,
	order_machine.OrderEvent_EventConfirm:       handlerConfirm,
	order_machine.OrderEvent_EventModify:        handlerModify,
	order_machine.OrderEvent_EventPay:           handlerPay,
	order_machine.OrderEvent_EventSend:          handlerSend,
	order_machine.OrderEvent_EventAccept:        handlerAccept,
	order_machine.OrderEvent_EventEvaluation:    handlerEvaluation,
	order_machine.OrderEvent_EventRefundAccept:  handlerRefundAccept,
	order_machine.OrderEvent_EventRefundCancel:  handlerRefundCancel,
	order_machine.OrderEvent_EventRefundConfirm: handlerRefundConfirm,
	order_machine.OrderEvent_EventRefundReject:  handlerRefundReject,
}
