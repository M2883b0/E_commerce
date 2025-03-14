package data

import (
	"auth_manage/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"os"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAuthRepo)

// Data .
type Data struct {
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	REDIS_ADDR := os.Getenv("REDIS_ADDR")
	if REDIS_ADDR == "" {
		REDIS_ADDR = "127.0.0.1:6379"
	}
	//redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	return &Data{rdb: rdb}, cleanup, nil
	//return &Data{}, cleanup, nil
}
