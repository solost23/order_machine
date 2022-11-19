package base

import (
	"context"

	"gorm.io/gorm"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/solost23/go_interface/gen_go/common"
)

type Action struct {
	ctx           context.Context
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	traceId       int64
	operator      int32
}

func (a *Action) SetContext(ctx context.Context) {
	a.ctx = ctx
	return
}

func (a *Action) SetHeader(header *common.RequestHeader) {
	a.traceId = header.TraceId
	a.operator = header.OperatorUid
	return
}

func (a *Action) SetMysql(mysqlConn *gorm.DB) {
	a.mysqlConnect = mysqlConn.WithContext(a.ctx)
	return
}

func (a *Action) SetkafkaProducer(kafkaProducer sarama.SyncProducer) {
	a.kafkaProducer = kafkaProducer
	return
}

func (a *Action) GetTraceId() int64 {
	return a.traceId
}

func (a *Action) GetOperator() int32 {
	return a.operator
}

func (this *Action) GetMysqlConnect() *gorm.DB {
	return this.mysqlConnect
}

func (this *Action) GetRedisClient() *redis.Client {
	return this.redisClient
}

func (this *Action) GetKafkaProducer() sarama.SyncProducer {
	return this.kafkaProducer
}

func (*Action) BuildError(code int32, msg string, header *common.RequestHeader) *common.ErrorInfo {
	if header == nil {
		header = new(common.RequestHeader)
	}
	return &common.ErrorInfo{
		Requester:   header.Requester,
		OperatorUid: header.OperatorUid,
		TraceId:     header.TraceId,
		Code:        code,
		Msg:         msg,
	}
}
