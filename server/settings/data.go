package settings

import (
	"github.com/cjburchell/tools-go/env"
)

var DataServiceAddress = env.Get("DATA_SERVICE", "http://localhost")
var DataServiceToken = env.Get("DATA_SERVICE_TOKEN", "token")
var PubSubAddress = env.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222")
var PubSubToken = env.Get("PUB_SUB_TOKEN", "token")
var MongoUrl = env.Get("MONGO_URL", "localhost")
