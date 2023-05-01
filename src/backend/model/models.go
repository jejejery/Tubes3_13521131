package model

type QnA struct {
	Id       uint   `gorm:"primaryKey" json:"Id"`
	Question string `gorm:"type: LONGTEXT" json:"Question"`
	Answer   string `gorm:"type: LONGTEXT" json:"Answer"`
}

type History struct {
	Id     uint   `gorm:"primaryKey" json:"Id"`
	Urutan uint   `gorm:"primaryKey" json:"Urutan"`
	Text   string `gorm:"type:LONGTEXT" json:"Text"`
}

type InputUser struct {
	Id        uint   `json:"Id"`
	InputText string `json:"InputText"`
	Algorithm bool   `json:"Algorithm"` //true : kmp, false : bm
	Answer    string `json:"Answer"`
}
