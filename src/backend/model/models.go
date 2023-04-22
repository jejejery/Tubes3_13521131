package model

type QnA struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Question string `gorm: "type: LONGTEXT" json:"question"`
	Answer   string `gorm: "type: LONGTEXT" json:"answer"`
}

type History struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Urutan uint   `gorm:"primaryKey" json:"urutan"`
	Text   string `gorm: "type: LONGTEXT" json:"text"`
}
