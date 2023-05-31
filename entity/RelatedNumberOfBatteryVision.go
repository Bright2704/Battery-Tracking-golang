package entity

type RelatedNumber struct {
	ID uint `json: "id" bson: "id"`

	RelatedProcess RelatedProcess `json: "relatedProcess" bson: "relatedProcess"`
}