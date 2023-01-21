package create_order

import (
	"context"
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/common"
	"github.com/solost23/protopb/gen/go/protos/order_machine"
	"order_machine/configs"
	"order_machine/internal/models"
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
	action := Action{}
	action.SetMysql(mdb)

	type test struct {
		ctx     context.Context
		request *order_machine.CreateOrderRequest
		err     error
	}

	tests := []test{
		{
			ctx: context.Background(),
			request: &order_machine.CreateOrderRequest{
				Header: &common.RequestHeader{
					OperatorUid: 1,
					TraceId:     34343,
				},
				CourseIds:   []uint32{1, 2},
				CourseNums:  []uint32{10, 20},
				OrderSource: order_machine.OrderSource_App,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		reply, err := action.Deal(test.ctx, test.request)
		if err != test.err {
			t.Error(err)
		}
		fmt.Println(reply.OrderId)
	}
}
