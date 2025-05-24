package redis

import (
	"context"
	"encoding/json"
	"fmt"
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

// DeleteEvent implements repoiface.Redis.
func (c *RedisClient) DeleteEvent(ctx context.Context,dbFrom int, indexKey string, dataKey string) (err error) {
	var (
		cliData, ok = c.cliMap[c.dbIndex]
	)
	if !ok {
		return fmt.Errorf("error db index")
	}

	// Del index
	err = c.cliIndex.Client().Del(ctx, indexKey).Err()
	if err != nil {
		return err
	}

	// Del data
	if err := cliData.Client().Del(ctx, dataKey).Err(); err != nil {
		return err
	}

	return nil
}

// SetEvent implements repoiface.Redis.
func (c *RedisClient) SetEvent(ctx context.Context, event model.Event, indexKey string, triggerKey string, dataKey string, releaseEvent time.Duration) (err error) {
	var (
		dbIndex     = c.IncKeyIndex(ctx)
		cliData, ok = c.cliMap[c.dbIndex]
		eventByte   []byte
	)
	if !ok {
		return fmt.Errorf("error db index")
	}

	if eventByte, err = json.Marshal(event); err != nil {
		return err
	}

	// set index
	ok, err = c.cliIndex.Client().SetNX(ctx, indexKey, dbIndex, releaseEvent).Result()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("event already exist")
	}
	// setTrigger
	if err := cliData.Client().Set(ctx, triggerKey, "trigger event", releaseEvent).Err(); err != nil {
		return err
	}

	if err := cliData.Client().Set(ctx, dataKey, eventByte, releaseEvent+(5*time.Minute)).Err(); err != nil {
		return err
	}

	return nil
}

// LockEventFromDb implements repoiface.Redis.
func (c *RedisClient) LockEventFromDb(ctx context.Context, dbFrom int, lockKey string) error {
	var (
		cliData, ok = c.cliMap[dbFrom]
	)
	if !ok {
		return fmt.Errorf("error db index")
	}
	return cliData.Client().SetNX(ctx, lockKey, "on-process-call", time.Second*5).Err()
}

// LockEventFromDb implements repoiface.Redis.
func (c *RedisClient) UnlockEventFromDb(ctx context.Context, dbFrom int, lockKey string) error {
	var (
		cliData, ok = c.cliMap[dbFrom]
	)
	if !ok {
		return fmt.Errorf("error db index")
	}
	return cliData.Client().Del(ctx, lockKey).Err()
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

func (c *RedisClient) IncKeyIndex(ctx context.Context) int32 {
	if c.dbIndex >= c.maxDbUse {
		c.dbIndex = 1
	} else {
			c.dbIndex++
	}
	return int32(c.dbIndex)
}

func (c *RedisClient) GetListOfListener(ctx context.Context) map[int]*redis.PubSub {
	mapOfListener := map[int]*redis.PubSub{}
	for k, cli := range c.cliMap {
		mapOfListener[k] = cli.ListenEvent(ctx, strconv.Itoa(k), cRedis.EventExpired)
	}
	return mapOfListener
}
