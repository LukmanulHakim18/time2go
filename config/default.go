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

	"pubsub_emulator_host_port": "",
	"pubsub_credential":         "",
	"pubsub_project_id":         "",
	"pod_name":                  "unknown pod",
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
