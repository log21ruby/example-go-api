package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ctx = context.Background()

// RedisConnection init
func RedisConnection(db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password:     viper.GetString("redis.password"),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		DB:           db,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Fatalf("cannot open redis connection : %s", err)
	}

	return rdb
}
