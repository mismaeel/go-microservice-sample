package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

)

type (
	Movie struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Title     string        `validate:"min=3,max=50,regexp=[a-zA-Z]" json:"title"`
		Rating    float32       `json:"rating"`
		Director  string        `json:"director"`
		Actors    []string      `json:"actors"`
		CreatedOn time.Time     `json:"createdAt,omitempty"`
	}
)
