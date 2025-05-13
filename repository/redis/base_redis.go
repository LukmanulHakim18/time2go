package redis

import (
	"context"
	"fmt"
	"net/url"

	cRedis "github.com/LukmanulHakim18/core/redis"
	"github.com/LukmanulHakim18/time2go/config/logger"
	"github.com/LukmanulHakim18/time2go/repository"
)

type RedisConfig struct {
	host        string
	password    string
	dbUse       int
	port        int
	checkHealth bool
}

func NewRedisConfig(host string, port int, password string, dbUse int, checkHealth bool) repository.RepoConf {
	return &RedisConfig{
		host:        host,
		port:        port,
		dbUse:       dbUse,
		password:    password,
		checkHealth: checkHealth,
	}
}

func (conf *RedisConfig) Init(r *repository.Repository) error {
	if conf.host == "" {
		return fmt.Errorf("redis repository host cannot be empty")
	}
	if conf.port == 0 {
		return fmt.Errorf("redis repository port cannot be empty")
	}
	_, err := url.Parse(fmt.Sprintf("%s:%d", conf.host, conf.port))
	if err != nil {
		return err
	}
	// init redis client
	cli := cRedis.NewRedis(conf.host, conf.port, conf.password, 0)
	errKea := cli.SetConfigKEA()
	if errKea != nil {
		logger.GetLogger().Error("error")
	}
	redisClient := &RedisClient{
		cliIndex: cli,
		cliMap:   make(map[int]cRedis.ClientRedis, conf.dbUse),
	}

	// init db use for optimization
	for i := 1; i <= conf.dbUse; i++ {
		cli := cRedis.NewRedis(conf.host, conf.port, conf.password, 1)
		redisClient.cliMap[i] = cli

	}
	if conf.checkHealth {
		if err := redisClient.HealthCheck(context.Background()); err != nil {
			return fmt.Errorf("redis repository failed to be created: healthCheck error, %s", err.Error())
		}
	}
	r.Redis = redisClient
	return nil
}

func (conf *RedisConfig) GetRepoName() string {
	return "Redis"
}
