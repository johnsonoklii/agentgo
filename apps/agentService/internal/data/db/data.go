package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/conf"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	slog "log"
	"os"
	"time"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewAgentRepo,
	NewAgentWorkspaceRepo,
	NewAgentVersionRepo,
	NewProviderRepo,
	NewModalRepo,
)

type Data struct {
	DB     *gorm.DB
	BizRDB *redis.Client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)

	db, err := newDB(c)
	if err != nil {
		return nil, nil, err
	}

	rdb, err := newBizRDB(c)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		helper.Info("closing data resources")
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
		_ = rdb.Close()
	}

	return &Data{
		DB:     db,
		BizRDB: rdb,
	}, cleanup, nil
}

func newDB(c *conf.Data) (*gorm.DB, error) {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("connect mysql failed: %w", err)
	}

	// 设置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func newBizRDB(c *conf.Data) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.BizRedis.Addr,
		Password: c.BizRedis.Password,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("connect redis failed: %w", err)
	}

	return rdb, nil
}
