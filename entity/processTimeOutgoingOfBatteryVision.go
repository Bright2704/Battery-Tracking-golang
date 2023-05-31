package entity

import "time"

type ProcessTimeOut struct {
	ID                   uint `json: "id" bson: "id"`
	ProcessTimeOutgoing1 time.Time `json: "processTimeOutgoing1" bson: "processTimeOutgoing1"`
	processTimeOutgoing2 time.Time `json: "processTimeOutgoing2" bson: "processTimeOutgoing2"`
	ProcessTimeOutgoing3 time.Time `json: "processTimeOutgoing3" bson: "processTimeOutgoing3"`
	processTimeOutgoing4 time.Time `json: "processTimeOutgoing4" bson: "processTimeOutgoing4"`
	ProcessTimeOutgoing5 time.Time `json: "processTimeOutgoing5" bson: "processTimeOutgoing5"`
	
}