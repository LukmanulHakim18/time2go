package redis

import (
	"context"
	"strconv"

	cRedis "github.com/LukmanulHakim18/core/redis"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	dbIndex  int
	maxDbUse int
	cliIndex cRedis.ClientRedis
	cliMap   map[int]cRedis.ClientRedis
}

func (c *RedisClient) HealthCheck(ctx context.Context) error {
	return c.cliIndex.Ping()
}

func (c *RedisClient) IncKeyIndex(ctx context.Context) {
	if c.dbIndex == c.maxDbUse {
		c.dbIndex = 1
	} else {
		c.dbIndex++
	}
}

func (c *RedisClient) GetListOfListener(ctx context.Context) map[int]*redis.PubSub {
	mapOfListener := map[int]*redis.PubSub{}
	for k, cli := range c.cliMap {
		mapOfListener[k] = cli.ListenEvent(ctx, strconv.Itoa(k), cRedis.EventExpired)
	}
	return mapOfListener
}
