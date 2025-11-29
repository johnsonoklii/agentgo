package data

import (
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/conf"
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	DTA  *Data
	RDB  *redis.Client
	once sync.Once
)

type Data struct {
	RDB *redis.Client
}

func GetData() *Data {
	return DTA
}

func InitData(c *conf.Data) {
	DTA, _ = NewData(c)
	if DTA == nil {
		panic("DTA is nil")
	}
}

func NewData(c *conf.Data) (*Data, error) {
	Init(c)
	return &Data{RDB: RDB}, nil
}

func Init(c *conf.Data) {
	once.Do(func() {
		InitRDB(c)
	})

	if RDB == nil {
		panic("RDB is nil")
	}
}

func InitRDB(c *conf.Data) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
	})
}
