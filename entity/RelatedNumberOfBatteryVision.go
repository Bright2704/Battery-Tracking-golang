package entity

type RelatedNumber struct {
	//ID             uint `gorm:"primaryKey;autoIncrement"`
	RelatedNumber string `json: "relatedNumber" `
	RelatedProcess RelatedProcess `json: "relatedProcess" `
}