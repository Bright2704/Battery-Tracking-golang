package entity

type RelatedNumber struct {
	//ID             uint `gorm:"primaryKey;autoIncrement"`
	RelatedNumber string `json: "relatedNumber" bson: "relatedNumber"`
	RelatedProcess RelatedProcess `json: "relatedProcess" bson: "relatedProcess` 
}