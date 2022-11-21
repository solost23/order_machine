package list_order

import (
	"context"
	"github.com/solost23/go_interface/gen_go/order_machine"
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

func (a *Action) Deal(ctx context.Context, request *order_machine.ListOrderRequest) (reply *order_machine.ListOrderResponse, err error) {
	db := a.GetMysqlConnect()

	sqlOrders, err := (&models.Order{}).WhereAll(db, "1 = ?", 1)
	if err != nil {
		return nil, err
	}
	records := make([]*order_machine.DetailOrder, 0, len(sqlOrders))
	for _, sqlOrder := range sqlOrders {
		events := make([]order_machine.OrderEvent, 0)
		for _, event := range fsm.StatusEvent[order_machine.OrderStatus(sqlOrder.Status)] {
			events = append(events, event)
		}
		records = append(records, &order_machine.DetailOrder{
			OrderId:     uint32(sqlOrder.ID),
			OrderStatus: uint32(sqlOrder.Status),
			OrderEvents: events,
		})
	}
	reply = &order_machine.ListOrderResponse{
		Records: records,
	}
	return reply, nil
}
