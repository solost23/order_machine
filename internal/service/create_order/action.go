package create_order

import (
	"context"
	"github.com/solost23/go_interface/gen_go/order_machine"
	"order_machine/internal/service/base"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (*Action) Deal(ctx context.Context, request *order_machine.CreateOrderRequest) (reply *order_machine.CreateOrderResponse, err error) {
	return nil, nil
}
