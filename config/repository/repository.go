package repository

import (
	"log"

	"github.com/LukmanulHakim18/time2go/config"
	"github.com/LukmanulHakim18/time2go/repository"
	"github.com/LukmanulHakim18/time2go/repository/redis"
)

var repo *repository.Repository

func LoadRepository() {
	repoList, err := repository.NewRepository([]repository.RepoConf{
		// TODO: add repository initialization here
		redis.NewRedisConfig(
			config.GetConfig("redis_host").GetString(),
			int(config.GetConfig("redis_port").GetInt()),
			config.GetConfig("redis_password").GetString(),
			int(config.GetConfig("db_use").GetInt()),
			config.GetConfig("check_healthy_repo").GetBool(),
		),
	})
	if err != nil {
		log.Fatalf("cannot initiate repository, with error: %v", err)
	}
	repo = repoList
}

func GetRepo() *repository.Repository {
	return repo
}
