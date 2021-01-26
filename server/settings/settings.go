package settings

import (
	"github.com/cjburchell/settings-go"
)

func Get(settings settings.ISettings) Config {
	return Config{
		SendOnReminder:   settings.GetBool("ALERT_REMINDER_ENABLE", false),
		RedisAddress:     settings.Get("REDIS_ADDRESS", "redis:6379"),
		RedisPassword:    settings.Get("REDIS_PASSWORD", ""),
		PubSubAddress:    settings.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222"),
		PubSubToken:      settings.Get("PUB_SUB_TOKEN", "token"),
		SlackDestination: settings.Get("SLACK_DESTINATION", ""),
		MongoUrl:         settings.Get("MONGO_URL", "localhost"),
		DataServiceToken: settings.Get("DATA_SERVICE_TOKEN", "token"),
	}
}

type Config struct {
	SendOnReminder   bool
	RedisAddress     string
	RedisPassword    string
	PubSubAddress    string
	PubSubToken      string
	SlackDestination string
	MongoUrl         string
	DataServiceToken string
}
