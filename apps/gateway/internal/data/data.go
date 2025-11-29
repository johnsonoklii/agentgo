package data

import (
	"context"
	"fmt"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/conf"
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	once sync.Once
	data *Data
	err  error
)

type Data struct {
	RDB *redis.Client
}

func Init(c *conf.Data) (*Data, error) {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
		})

		// Ping 检测 Redis 是否正常
		if _, pingErr := rdb.Ping(context.Background()).Result(); pingErr != nil {
			err = fmt.Errorf("redis connect error: %w", pingErr)
			return
		}

		data = &Data{RDB: rdb}
	})

	return data, err
}

// Get 全局获取 Data 实例
func Get() *Data {
	return data
}
