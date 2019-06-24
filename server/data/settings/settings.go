package settings

import (
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
)

var DataServiceToken = env.Get("DATA_SERVICE_TOKEN", "token")
var RedisAddress = env.Get("REDIS_ADDRESS", "redis:6379")
var RedisPassword = env.Get("REDIS_PASSWORD", "")
var Log = log.CreateDefaultSettings()
