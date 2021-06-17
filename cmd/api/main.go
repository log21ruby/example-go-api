package main

import (
	"example-go-api/config"
	"example-go-api/pkg/driver"
	"example-go-api/pkg/routing"
	"example-go-api/repository"
	"example-go-api/service/user"
	"os"
	"os/signal"

	"github.com/go-redis/cache/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	config.InitViper("config")
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	var (
		dbw            = driver.MySQLWeb()
		dbmg, ctx, err = driver.MongoConnection(viper.GetString("mongo.dsn"))
		dbrd           = driver.RedisConnection(viper.GetInt("redis.DB"))
		appCache       = cache.New(&cache.Options{
			Redis: dbrd,
		})
		userRepo    = repository.NewUserRepository(dbmg, dbw)
		userService = user.NewService(user.RepoInterface{User: userRepo}, appCache)
		userHandler = user.NewHandler(userService)

		newFiber  = routing.InitFiber()
		f, router = newFiber.InitFiberMiddleware()
	)
	defer func() {
		if err = dbmg.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	V1 := router.Group("/v1")
	userV1 := V1.Group("/user")
	userV1.Post("/:id", userHandler.GetUserWithCache)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		logrus.Info("Gracefully shutting down...")
		_ = f.Shutdown()
	}()
	if err := f.Listen(":" + viper.GetString("app.port")); err != nil {
		logrus.Fatalf("shutting down the server : %s", err)
	}
}
