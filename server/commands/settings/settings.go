package settings

import (
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
)

var Log = log.CreateDefaultSettings()
var PubSubAddress = env.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222")
var PubSubToken = env.Get("PUB_SUB_TOKEN", "token")

var DataServiceToken = env.Get("COMMAND_TOKEN", "token")
