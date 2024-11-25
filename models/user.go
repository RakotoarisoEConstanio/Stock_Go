package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm : "Unique"`
	Password []byte `json:"-"`
}
