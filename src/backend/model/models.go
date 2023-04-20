package model

type QnA struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type History struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Urutan uint   `gorm:"primaryKey" json:"urutan"`
	Text   string `json:"text"`
}
