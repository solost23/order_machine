package handler

import (
	"github.com/solost23/protopb/gen/go/protos/order_machine"
	"order_machine/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::user service
	order_machine.RegisterOrderMachineServer(config.Server, service.NewOrderMachineService(config.Sl, config.MysqlConnect, config.RedisClient, config.KafkaProducer))
	return
}
