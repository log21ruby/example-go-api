package repository

import (
	"example-go-api/model"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Repository db type
type Repository struct {
	db *gorm.DB
	cl *mongo.Client
	rd *redis.Client
}

// RedisRepositorier interface
type RedisRepositorier interface {
	Set(key string, val interface{}, td time.Duration) error
	Get(prefix string) ([]string, error)
	DelKey(prefix string) (int64, error)
	GetPrefix(prefix string) (string, error)
	DelKeyPrefix(prefix []string) (int64, error)
}

type UserRepositorier interface {
	FindOneUserWithUID(uid int, field ...interface{}) (model.User, error)
}

func NewUserRepository(client *mongo.Client, gdb *gorm.DB) UserRepositorier {
	return &Repository{
		cl: client,
		db: gdb,
	}
}
