package data

import (
	"time"

	"github.com/cjburchell/reefstatus-history/model"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoData struct {
	session *mgo.Session
	db      *mgo.Database
}

func (m *mongoData) setup(address string) error {
	session, err := mgo.Dial(address)
	if err != nil {
		return errors.WithStack(err)
	}

	err = session.Ping()
	if err != nil {
		return errors.WithStack(err)
	}

	m.db = session.DB(dbName)
	m.session = session

	return nil
}

func (m mongoData) saveData(data model.Data, collection string) error {
	err := m.db.C(collection).Insert(data)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

const dbName = "ReefStatus"
const dayCollection = "dayHistory"
const weekCollection = "weekHistory"
const yearCollection = "yearHistory"

func (m mongoData) SaveDayData(data model.Data) error {
	return m.saveData(data, dayCollection)
}

func (m mongoData) SaveWeekData(data model.Data) error {
	return m.saveData(data, weekCollection)
}

func (m mongoData) SaveYearData(data model.Data) error {
	return m.saveData(data, yearCollection)
}

func (m mongoData) GetLastTimeYearDataWasSaved() (time.Time, error) {
	var data model.Data
	err := m.db.C(yearCollection).Find(nil).Sort("-time").One(&data)

	if err == mgo.ErrNotFound {
		return time.Time{}, nil
	}

	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	return data.Time, nil
}

func (m mongoData) GetLastTimeWeekDataWasSaved() (time.Time, error) {
	var data model.Data
	err := m.db.C(weekCollection).Find(nil).Sort("-time").One(&data)

	if err == mgo.ErrNotFound {
		return time.Time{}, nil
	}

	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	return data.Time, nil
}

func (m mongoData) getDataPoints(collection string, dataType string) ([]model.Data, error) {

	var data []model.Data
	err := m.db.C(collection).Find(bson.M{"type": dataType}).All(&data)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (m mongoData) GetDayDataPoints(dataType string) ([]model.Data, error) {

	return m.getDataPoints(dayCollection, dataType)
}

func (m mongoData) GetYearDataPoints(dataType string) ([]model.Data, error) {
	return m.getDataPoints(yearCollection, dataType)
}

func (m mongoData) GetWeekDataPoints(dataType string) ([]model.Data, error) {
	return m.getDataPoints(weekCollection, dataType)
}

func (m mongoData) GetDataPointsFromLastHour(dataType string) ([]model.Data, error) {
	cutoffTime := time.Now().Add(-time.Hour)
	var data []model.Data
	err := m.db.C(dayCollection).Find(bson.M{"type": dataType, "time": bson.M{"$gt": cutoffTime}}).All(&data)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
