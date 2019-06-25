package state

import (
	"strconv"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoData struct {
	session *mgo.Session
	c       *mgo.Collection
}

type StateData interface {
	UpdateAlarmSent(sent bool) (bool, error)
	UpdateReminderSent(index int, sent bool) (bool, error)
}

// New History data
func New(address string) (StateData, error) {
	data := mongoData{}
	err := data.setup(address)
	return data, err
}

const dbName = "ReefStatus"
const stateCollection = "state"

func (m *mongoData) setup(address string) error {
	session, err := mgo.Dial(address)
	if err != nil {
		return errors.WithStack(err)
	}

	err = session.Ping()
	if err != nil {
		return errors.WithStack(err)
	}

	m.session = session
	m.c = session.DB(dbName).C(stateCollection)

	return nil
}

type stateItem struct {
	ID    string `bson:"_id"`
	State bool   `bson:"state"`
}

// UpdateAlarmSent will update the alarm sent state and return true if we updated it
func (m mongoData) UpdateAlarmSent(sent bool) (bool, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"state": sent}},
		Upsert:    true,
		ReturnNew: false,
	}

	var item stateItem
	_, err := m.c.Find(bson.M{"_id": "alarm"}).Apply(change, &item)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return item.State != sent, nil
}

func (m mongoData) UpdateReminderSent(index int, sent bool) (bool, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"state": sent}},
		Upsert:    true,
		ReturnNew: false,
	}

	var item stateItem
	_, err := m.c.Find(bson.M{"_id": strconv.Itoa(index)}).Apply(change, &item)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return item.State != sent, nil
}
