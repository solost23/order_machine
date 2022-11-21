package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/solost23/go_interface/gen_go/order_machine"
	"gorm.io/gorm"
	"order_machine/internal/service/create_order"
	"order_machine/internal/service/list_order"
	"order_machine/internal/service/switch_order_status"
)

type OrderMachineService struct {
	mdb           *gorm.DB
	rdb           *redis.Client
	kafkaProducer sarama.SyncProducer
	order_machine.UnimplementedOrderMachineServer
}

func NewOrderMachineService(mdb *gorm.DB, rdb *redis.Client, kafkaProducer sarama.SyncProducer) *OrderMachineService {
	return &OrderMachineService{
		mdb:           mdb,
		rdb:           rdb,
		kafkaProducer: kafkaProducer,
	}
}

// create order
func (h *OrderMachineService) CreateOrder(ctx context.Context, request *order_machine.CreateOrderRequest) (reply *order_machine.CreateOrderResponse, err error) {
	action := create_order.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mdb)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

func (h *OrderMachineService) ListOrder(ctx context.Context, request *order_machine.ListOrderRequest) (reply *order_machine.ListOrderResponse, err error) {
	action := list_order.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mdb)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

func (h *OrderMachineService) SwitchOrderState(ctx context.Context, request *order_machine.SwitchOrderStateRequest) (reply *order_machine.SwitchOrderStateResponse, err error) {
	action := switch_order_status.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mdb)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}
