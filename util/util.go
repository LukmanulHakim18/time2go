package util

import (
	"strings"

	"github.com/LukmanulHakim18/core/redis"
	"github.com/LukmanulHakim18/time2go/config"
)

type KeyType string

const (
	KEY_TYPE_INDEX   KeyType = "index"
	KEY_TYPE_TRIGGER KeyType = "trigger"
	KEY_TYPE_DATA    KeyType = "data"
	KEY_TYPE_LOCK    KeyType = "lock"
)

// time2go:{keyType}:clientName:eventName:eventId-14
func CreateEventKey(keyType KeyType, clientName, eventName, identifier string) string {
	return redis.BuildKey(config.GetConfig("app_name").GetString(), string(keyType), clientName, eventName, identifier)
}

// time2go:trigger:clientName:eventName:eventId-14
func CheckIsEventKey(key string) bool {
	sliceKey := strings.Split(key, ":")

	if len(sliceKey) != 5 {
		return false
	}
	if sliceKey[0] != config.GetConfig("app_name").GetString() {
		return false
	}
	if sliceKey[1] != string(KEY_TYPE_TRIGGER) {
		return false
	}
	return true
}

// in : time2go:trigger:clientName:eventName:eventId-14
// out : time2go:data:clientName:eventName:eventId-14
func GetDataKeyFromEventKey(eventKey string) string {
	if !CheckIsEventKey(eventKey) {
		return ""
	}
	sliceEventKey := strings.Split(eventKey, ":")
	sliceEventKey[1] = string(KEY_TYPE_DATA)
	return redis.BuildKey(sliceEventKey...)

}
