package models

type User struct {
	Name         string        `db:"name" json:"user"`
	Age          int           `db:"age" json:"age"`
	Email        string        `db:"email" json:"email"`
	UserPassword *UserPassword `db:"user_password" json:"user_password"`
	AddressList  []Address     `json:"address_list"`

	BaseEntity
}
