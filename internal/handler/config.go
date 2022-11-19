package handler

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	Server        *grpc.Server
	MysqlConnect  *gorm.DB
	RedisClient   *redis.Client
	KafkaProducer sarama.SyncProducer
}
