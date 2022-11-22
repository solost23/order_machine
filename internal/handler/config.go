package handler

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	Server        *grpc.Server
	Sl            *slog.SugaredLogger
	MysqlConnect  *gorm.DB
	RedisClient   *redis.Client
	KafkaProducer sarama.SyncProducer
}
