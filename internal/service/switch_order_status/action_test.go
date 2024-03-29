package switch_order_status

import (
	"context"
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/common"
	"github.com/solost23/protopb/gen/go/protos/order_machine"
	"order_machine/configs"
	"order_machine/internal/models"
	"order_machine/pkg/helper"
	"testing"
)

func TestAction_Deal(t *testing.T) {
	mysqlConf := &configs.MySQLConf{
		DataSourceName:  "root:123@tcp(127.0.0.1:3306)/order_machine?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
		MaxIdleConn:     20,
		MaxOpenConn:     10,
		MaxConnLifeTime: 100,
	}
	mdb, _ := models.InitMysql(mysqlConf)
	sl, _ := helper.InitLogger("../../../", "logs", "debug")
	action := Action{}
	action.SetMysql(mdb)
	action.SetSl(sl)

	type test struct {
		ctx     context.Context
		request *order_machine.SwitchOrderStateRequest
		err     error
	}

	tests := []test{
		{
			ctx: context.Background(),
			request: &order_machine.SwitchOrderStateRequest{
				Header: &common.RequestHeader{
					OperatorUid: 1,
					TraceId:     34344,
				},
				OrderId:    2,
				OrderEvent: order_machine.OrderEvent_EventRefundConfirm,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		reply, err := action.Deal(test.ctx, test.request)
		if err != test.err {
			t.Error(err)
		}
		if reply != nil {
			fmt.Println(reply.OrderStatus)
		}
	}
}
