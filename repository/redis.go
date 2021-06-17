package repository

import (
	"context"
	"time"
)

// Set repo
func (repo *Repository) Set(key string, val interface{}, td time.Duration) error {
	var ctx = context.Background()
	err := repo.rd.Set(ctx, key, val, td).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetKeyPrefix repo
func (repo *Repository) GetKeyPrefix(prefix string) ([]string, error) {
	var ctx = context.Background()
	key, _ := repo.rd.Keys(ctx, prefix).Result()
	return key, nil
}

// GetPrefix repo
func (repo *Repository) Get(prefix string) (string, error) {
	var ctx = context.Background()
	resp, err := repo.rd.Get(ctx, prefix).Result()
	if err != nil {
		return "", err
	}
	return resp, nil
}

// DelKey repo
func (repo *Repository) DelKey(prefix string) (int64, error) {
	var ctx = context.Background()
	resp, err := repo.rd.Del(ctx, prefix).Result()
	if err != nil {
		return 0, err
	}
	return resp, nil
}

// DelKeyPrefix repo
func (repo *Repository) DelKeyPrefix(prefix []string) (int64, error) {
	var ctx = context.Background()
	resp, err := repo.rd.Del(ctx, prefix...).Result()
	if err != nil {
		return 0, err
	}
	return resp, nil
}
