package list_order

import (
	"context"
	"fmt"
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
		request *order_machine.ListOrderRequest
		err     error
	}

	tests := []test{
		{
			ctx:     context.Background(),
			request: nil,
			err:     nil,
		},
	}

	for _, test := range tests {
		reply, err := action.Deal(test.ctx, test.request)
		if err != nil {
			t.Error(err)
		}
		for _, record := range reply.Records {
			fmt.Printf("id: %v, status: %v, events: %v", record.GetOrderId(), record.GetOrderStatus(), record.GetOrderEvents())
		}
	}
}
