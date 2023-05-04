package model

type QnA struct {
	Id       uint   `gorm:"primaryKey" json:"Id"`
	Question string `gorm:"type: LONGTEXT" json:"Question"`
	Answer   string `gorm:"type: LONGTEXT" json:"Answer"`
}

type InputUser struct {
	Id        uint   `json:"Id"`
	Session   int64  `json:"Session" gorm:"foreignKey:Sessions"`
	InputText string `json:"InputText"`
	Algorithm bool   `json:"Algorithm"` //true : kmp, false : bm
	Answer    string `json:"Answer"`
}

type Sessions struct {
	Id      uint  `json:"Id"`
	Session int64 `json:"Session"`
}
