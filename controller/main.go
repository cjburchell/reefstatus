package main

import (
	"strconv"
	"time"

	logSettings "github.com/cjburchell/uatu-go/settings"

	"github.com/cjburchell/settings-go"
	"github.com/cjburchell/tools-go/env"

	appSettings "github.com/cjburchell/reefstatus/controller/settings"

	logger "github.com/cjburchell/uatu-go"

	"github.com/cjburchell/reefstatus/controller/commands"

	"github.com/cjburchell/reefstatus/controller/update"

	"github.com/cjburchell/reefstatus/common/communication"
	"github.com/cjburchell/reefstatus/controller/service"
)

const logRate = time.Second * 30

func Update(session communication.PublishSession, isInitial bool) error {
	return session.Publish(communication.UpdateMessage, strconv.FormatBool(isInitial))
}

func main() {
	set := settings.Get(env.Get("ConfigFile", ""))
	log := logger.Create(logSettings.Get(set.GetSection("Logging")))
	appConfig := appSettings.Get(set)

	controller, err := service.NewController(appConfig.DataServiceAddress, appConfig.DataServiceToken)
	if err != nil {
		log.Fatal(err, "Unable to Connect to data database:")
	}

	session, err := communication.NewSession(appConfig.PubSubAddress, appConfig.PubSubToken, log)
	if err != nil {
		log.Fatal(err, "Unable to Connect to pub sub")
	}

	defer session.Close()
	go commands.Handle(session, controller, log, appConfig.Connection)

	for {
		err = update.All(controller, log, appConfig.Connection)
		if err == nil {
			err = Update(session, true)
			if err != nil {
				log.Errorf(err, "Unable to send first update")
			}
			break
		}

		log.Error(err, "Unable to do first update")
		log.Debugf("RefreshSettings Sleeping for %s", logRate.String())
		<-time.After(logRate)
		continue
	}

	updateCount := 0
	for {
		log.Debugf("RefreshSettings Sleeping for %s", logRate.String())
		<-time.After(logRate)
		if updateCount%20 == 19 {
			err = update.All(controller, log, appConfig.Connection)
			if err != nil {
				log.Errorf(err, "Unable to update")
			}

		} else {
			err = update.State(controller, log, appConfig.Connection)
			if err != nil {
				log.Errorf(err, "Unable to update state")
			}
		}

		err = Update(session, false)
		if err != nil {
			log.Errorf(err, "Unable to send update")
		}
		updateCount++
	}
}
