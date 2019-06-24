package state

import (
	"strconv"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session
var c *mgo.Collection

const dbName = "ReefStatus"
const stateCollection = "state"

func Setup(url string) error {
	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		return errors.Wrap(errors.WithStack(err), "Settings: Error connecting to Mongo")
	}

	err = session.Ping()
	if err != nil {
		return errors.Wrap(errors.WithStack(err), "Settings: Unable to ping mongo")
	}

	c = session.DB(dbName).C(stateCollection)
	return nil
}

type stateItem struct {
	ID    string `bson:"_id"`
	State bool   `bson:"state"`
}

// UpdateAlarmSent will update the alarm sent state and return true if we updated it
func UpdateAlarmSent(sent bool) (bool, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"state": sent}},
		Upsert:    true,
		ReturnNew: false,
	}

	var item stateItem
	_, err := c.Find(bson.M{"_id": "alarm"}).Apply(change, &item)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return item.State != sent, nil
}

func UpdateReminderSent(index int, sent bool) (bool, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"state": sent}},
		Upsert:    true,
		ReturnNew: false,
	}

	var item stateItem
	_, err := c.Find(bson.M{"_id": strconv.Itoa(index)}).Apply(change, &item)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return item.State != sent, nil
}
