package switch_order_status

import (
	"context"
	"errors"
	"github.com/solost23/protopb/gen/go/protos/order_machine"
	"gorm.io/gorm"
	"order_machine/internal/models"
	"order_machine/internal/service/base"
	"order_machine/pkg/fsm"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *order_machine.SwitchOrderStateRequest) (reply *order_machine.SwitchOrderStateResponse, err error) {
	db := a.GetMysqlConnect()

	sqlOrder, err := (&models.Order{}).WhereOne(db, "id = ?", request.GetOrderId())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		a.GetSl().Infof("订单[%d]不存在，无法进行订单状态变更", request.GetOrderId())
		return nil, errors.New("此订单不存在，无法进行订单状态变更")
	}

	orderMachine, err := fsm.NewFSM(order_machine.OrderStatus(sqlOrder.Status))
	if err != nil {
		return nil, err
	}
	sqlOrder.RefundAcceptIntroduce = request.GetRefundAcceptIntroduce()
	sqlOrder.Evaluation = request.GetEvaluation()
	orderStatus, err := orderMachine.Call(request.GetOrderEvent(), fsm.WithDB(db), fsm.WithOrderInfo(sqlOrder))
	if err != nil {
		a.GetSl().Info(err.Error())
		return nil, err
	}
	// 修改数据库内订单状态
	value := map[string]interface{}{
		"status": orderStatus,
	}
	err = (&models.Order{}).Updates(db, value, "id = ?", request.GetOrderId())
	if err != nil {
		a.GetSl().Error(err.Error())
		return nil, err
	}
	reply = &order_machine.SwitchOrderStateResponse{
		OrderId:     uint32(sqlOrder.ID),
		OrderStatus: orderStatus,
	}
	return reply, nil
}
