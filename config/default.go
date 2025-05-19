package config

import (
	"strings"

	commonConfig "github.com/LukmanulHakim18/core/config"
)

var appConfig map[string]commonConfig.Value

var defaultConfig = map[string]interface{}{
	"app_name":      "time2go",
	"grpc_port":     6000,
	"rest_port":     8000,
	"log_level":     "info",
	"log_directory": "",
	"check_healthy_repo": true,

	"pubsub_emulator_host_port": "",
	"pubsub_credential":         "",
	"pubsub_project_id":         "",
	"pod_name":                  "unknown pod",

	"redis_host":         "",
	"redis_port":         0,
	"redis_password":     "",
	"db_use":             5,
	
	"worker_count": 10,
	"queue_size":   100,

	"retry_time_delay": "5s",
}

func LoadConfigMap() {
	appConfig = commonConfig.LoadConfig(defaultConfig)
}

func GetConfig(key string) (val commonConfig.Value) {
	if v, ok := appConfig[strings.ToLower(key)]; ok {
		val = v
	}
	return
}
