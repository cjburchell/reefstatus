package settings

import (
	"github.com/cjburchell/reefstatus/controller/profilux"
	"github.com/cjburchell/tools-go/env"
)

func newConnectionSettings() (connection profilux.Settings) {
	connection.Address = env.Get("PROFILUX_ADDRESS", "192.168.3.10")
	connection.Port = env.GetInt("PROFILUX_PORT", 80)
	connection.Protocol = env.Get("PROFILUX_PROTOCOL", profilux.ProtocolHTTP)
	connection.ControllerAddress = 1
	return
}

// Connection settings
var Connection = newConnectionSettings()
var PubSubAddress = env.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222")
var PubSubToken = env.Get("PUB_SUB_TOKEN", "token")
var DataServiceAddress = env.Get("DATA_SERVICE", "http://localhost")
var DataServiceToken = env.Get("DATA_SERVICE_TOKEN", "token")
