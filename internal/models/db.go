package models

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"order_machine/configs"
	"os"
	"time"
)

func InitMysql(mysqlConf *configs.MySQLConf) (db *gorm.DB, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlConf.DataSourceName,
		DefaultStringSize: 100,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// 引入读写分离
	_ = db.Use(dbresolver.Register(dbresolver.Config{
		// master为sources, master作为replicas, 暂时读写都设置为一个数据库
		Sources: []gorm.Dialector{mysql.New(mysql.Config{
			DSN:               mysqlConf.DataSourceName,
			DefaultStringSize: 100,
		})},
		Replicas: []gorm.Dialector{mysql.New(mysql.Config{
			DSN:               mysqlConf.DataSourceName,
			DefaultStringSize: 100,
		})},
	}))

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Panic(err)
	}
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConn)
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxConnLifeTime) * time.Second)

	if err != nil {
		return nil, err
	}
	return db, nil
}
