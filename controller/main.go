package main

import (
	"time"

	"github.com/cjburchell/reefstatus/controller/settings"

	"github.com/cjburchell/go-uatu"
	logSettings "github.com/cjburchell/go-uatu/settings"

	"github.com/cjburchell/reefstatus/controller/commands"

	"github.com/cjburchell/reefstatus/controller/update"

	"github.com/cjburchell/reefstatus/common/communication"
	"github.com/cjburchell/reefstatus/common/data"
)

const logRate = time.Second * 30

func main() {
	err := logSettings.SetupLogger()
	if err != nil {
		log.Warn("Unable to Connect to logger")
	}

	controller, err := data.NewController(settings.DataServiceAddress, settings.DataServiceToken)
	if err != nil {
		log.Fatal(err, "Unable to Connect to data database:")
	}

	session, err := communication.NewSession(settings.PubSubAddress, settings.PubSubToken)
	if err != nil {
		log.Fatal(err, "Unable to Connect to pub sub")
	}

	defer session.Close()
	go commands.Handle(session, controller)

	for {
		err = update.All(controller)
		if err == nil {
			err = communication.Update(session, true)
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
			err = update.All(controller)
			if err != nil {
				log.Errorf(err, "Unable to update")
			}

		} else {
			err = update.State(controller)
			if err != nil {
				log.Errorf(err, "Unable to update state")
			}
		}

		err = communication.Update(session, false)
		if err != nil {
			log.Errorf(err, "Unable to send update")
		}
		updateCount++
	}
}
