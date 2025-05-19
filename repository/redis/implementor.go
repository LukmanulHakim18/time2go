package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	cRedis "github.com/LukmanulHakim18/core/redis"
	"github.com/LukmanulHakim18/time2go/model"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	dbIndex  int
	maxDbUse int
	cliIndex cRedis.ClientRedis
	cliMap   map[int]cRedis.ClientRedis
}

// SetEvent implements repoiface.Redis.
func (c *RedisClient) SetEvent(ctx context.Context, event model.Event, indexKey string, triggerKey string, dataKey string, releaseEvent time.Duration) error {
	panic("unimplemented")
}

// DeleteFromDb implements repoiface.Redis.
func (c *RedisClient) DeleteFromDb(ctx context.Context, dbFrom int, dataKey string) error {
	panic("unimplemented")
}

// LockEventFromDb implements repoiface.Redis.
func (c *RedisClient) LockEventFromDb(ctx context.Context, dbFrom int, dataKey string) error {
	panic("unimplemented")
}

// GetDataFromDb implements repoiface.Redis.
func (c *RedisClient) GetDataFromDb(ctx context.Context, dbFrom int, dataKey string) (result model.Event, err error) {
	result = model.Event{}
	cli, ok := c.cliMap[dbFrom]
	if !ok {
	}
	resultByte, err := cli.Client().Get(ctx, dataKey).Bytes()
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(resultByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
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
