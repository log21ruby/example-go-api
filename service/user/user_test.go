package user

import (
	"example-go-api/config"
	"example-go-api/pkg/driver"
	"example-go-api/repository"
	"testing"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var _ = (func() interface{} {
	config.InitViper("../../config")
	return nil
}())

var (
	dbw        = driver.MySQLWeb()
	dbmg, _, _ = driver.MongoConnection(viper.GetString("mongo.dsn"))
	dbrd       = driver.RedisConnection(viper.GetInt("redis.DB"))
	appCache   = cache.New(&cache.Options{
		Redis: dbrd,
	})
	userRepo    = repository.NewUserRepository(dbmg, dbw)
	userService = NewService(RepoInterface{User: userRepo}, appCache)
)

func TestGetUserWithUIDCache(t *testing.T) {
	resp, err := userService.GetUserWithUIDCache(1*time.Second, "core:user:358254", 358254, []string{})
	assert.NotNil(t, resp, "this user %v", resp)
	assert.NoError(t, err, "this not error %v", err)
}
