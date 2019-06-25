package settings

import (
	"github.com/cjburchell/tools-go/env"
)

var RedisAddress = env.Get("REDIS_ADDRESS", "redis:6379")
var RedisPassword = env.Get("REDIS_PASSWORD", "")
var PubSubAddress = env.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222")
var PubSubToken = env.Get("PUB_SUB_TOKEN", "token")
var SlackDestination = env.Get("SLACK_DESTINATION", "")
var MongoUrl = env.Get("MONGO_URL", "localhost")
var DataServiceToken = env.Get("DATA_SERVICE_TOKEN", "token")
