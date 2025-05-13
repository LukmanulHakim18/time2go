package redis

import (
	"context"
	"fmt"
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

func (c *RedisClient) ListenEventInDb(ctx context.Context, dbNumb int) (*redis.PubSub, error) {
	client, ok := c.cliMap[dbNumb]
	if !ok {
		return nil, fmt.Errorf("error select db")
	}
	pubsub := client.ListenEvent(ctx, strconv.Itoa(dbNumb), cRedis.EventExpired)
	return pubsub, nil
}
