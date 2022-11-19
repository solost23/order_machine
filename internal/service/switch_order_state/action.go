package switch_order_state

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

func (*Action) Deal(ctx context.Context, request *order_machine.SwitchOrderStateRequest) (reply *order_machine.SwitchOrderStateResponse, err error) {
	return nil, nil
}
