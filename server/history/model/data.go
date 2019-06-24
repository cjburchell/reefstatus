package model

import "time"

// Data model
type Data struct {
	Time  time.Time `json:"time" bson:"time"`
	Type  string    `json:"type" bson:"type"`
	Value float64   `json:"value" bson:"value"`
}
