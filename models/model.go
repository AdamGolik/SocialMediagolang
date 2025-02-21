package models

type User struct {
	Id         uint32  `json:"id" gorm:"primary_key"`
	Nickname   string  `json:"nickname"`
	Password   string  `json:"password"`
	Images     []Image `gorm:"foreignKey:UserID"`
	ProfilePic *Image  `json:"profile_pic,omitempty" gorm:"foreignKey:UserID"`
}

type Image struct {
	Id     uint32 `json:"id" gorm:"primary_key"`
	UserID uint32 // dodaj pole klucza obcego
	Public bool   `json:"public"`
	Name   string `json:"name"`
}
