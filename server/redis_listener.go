package server

import (
	"context"

	"github.com/LukmanulHakim18/core/redis"
	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/repository"
)

type RedisListener struct {
	dbUsed   int
	listener map[int]redis.ClientRedis
}

func NewRedisListener(ctx context.Context, repo repository.Repository) *RedisListener {

	rl := &RedisListener{
		dbUsed: int(config.GetConfig("db_used").GetInt()),
	}
	for i := 1; i <= rl.dbUsed; i++ {

		
	}
	return rl
}

func (rl *RedisListener) Shutdown(ctx context.Context) error {
	return nil
}
