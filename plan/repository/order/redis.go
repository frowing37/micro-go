package order

import (
	"context"
	"github.com/dreamsofcode-io/orders-api/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepos struct {
	Client *redis.Client
}

func (r *RedisRepos) Insert(ctx context.Context, order model.Order) error {
	return nil
}
