package entity

import (
	"time"
)

type ProcessTimeIn struct {
	ID                   uint       `json:"id" bson:"id"`
	ProcessTimeInComing1 time.Time  `json:"processTimeInComing1" bson:"processTimeInComing1"`
	ProcessTimeInComing2 time.Time  `json:"processTimeInComing2" bson:"processTimeInComing2"`
	ProcessTimeInComing3 time.Time  `json:"processTimeInComing3" bson:"processTimeInComing3"`
	ProcessTimeInComing4 time.Time  `json:"processTimeInComing4" bson:"processTimeInComing4"`
	ProcessTimeInComing5 time.Time  `json:"processTimeInComing5" bson:"processTimeInComing5"`
}