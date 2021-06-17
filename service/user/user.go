package user

import (
	"example-go-api/model"
	"example-go-api/repository"
	"time"

	"github.com/go-redis/cache/v8"
)

type Servicer interface {
	GetUserWithUIDCache(ttl time.Duration, key string, uid int, field []string) (model.User, error)
}

type RepoInterface struct {
	User repository.UserRepositorier
}

type Service struct {
	userRepo repository.UserRepositorier
	Cache    *cache.Cache
}

func NewService(repo RepoInterface, cache *cache.Cache) *Service {
	return &Service{
		userRepo: repo.User,
		Cache:    cache,
	}
}

func (s *Service) GetUserWithUIDCache(ttl time.Duration, key string, uid int, field []string) (model.User, error) {
	var resp model.User
	err := s.Cache.Once(&cache.Item{
		Key:   key,
		TTL:   ttl,
		Value: &resp,
		Do: func(*cache.Item) (interface{}, error) {
			return s.userRepo.FindOneUserWithUID(uid, field)
		},
	})
	return resp, err
}
