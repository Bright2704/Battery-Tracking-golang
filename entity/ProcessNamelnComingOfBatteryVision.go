package entity

type ProcessNameIn struct {
	ID uint `json: "id" bson: "id"`

	ProcessNameIn1 bool `json: "processNameInComing1" bson: "processNameInComing1"`
	ProcessNameIn2 bool `json: "processNameInComing2" bson: "processNameInComing2"`
	ProcessNameIn3 bool `json: "processNameInComing3" bson: "processNameInComing3"`
	ProcessNameIn4 bool `json: "processNameInComing4" bson: "processNameInComing4"`
	ProcessNameIn5 bool `json: "processNameInComing5" bson: "processNameInComing5"`
}