package settings

import (
	"github.com/cjburchell/reefstatus/controller/profilux"
	"github.com/cjburchell/settings-go"
)

func Get(settings settings.ISettings) Config {
	return Config{
		Connection:         newConnectionSettings(settings),
		PubSubAddress:      settings.Get("PUB_SUB_ADDRESS", "tcp://localhost:4222"),
		PubSubToken:        settings.Get("PUB_SUB_TOKEN", "token"),
		DataServiceAddress: settings.Get("DATA_SERVICE", "http://localhost"),
		DataServiceToken:   settings.Get("DATA_SERVICE_TOKEN", "token"),
	}
}

func newConnectionSettings(settings settings.ISettings) (connection profilux.Settings) {
	connection.Address = settings.Get("PROFILUX_ADDRESS", "192.168.3.10")
	connection.Port = settings.GetInt("PROFILUX_PORT", 80)
	connection.Protocol = settings.Get("PROFILUX_PROTOCOL", profilux.ProtocolHTTP)
	connection.ControllerAddress = 1
	return
}

type Config struct {
	Connection         profilux.Settings
	PubSubAddress      string
	PubSubToken        string
	DataServiceAddress string
	DataServiceToken   string
}
