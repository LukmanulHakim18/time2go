package util

import (
	"github.com/LukmanulHakim18/core/redis"
	"github.com/LukmanulHakim18/time2go/config"
)

type KeyType string

func CreateEventKey(keyType KeyType, clientName, eventName, identifier string) string {
	return redis.BuildKey(config.GetConfig("app_name").GetString(), clientName, eventName, identifier, string(keyType))
}
