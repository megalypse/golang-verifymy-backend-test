package models

type User struct {
	Name         string        `json:"name" db:"name"`
	Age          int           `json:"age" db:"age"`
	Email        string        `json:"email" db:"email"`
	UserPassword *UserPassword `json:"user_password" db:"user_password"`
	AddressList  []Address     `json:"address_list"`

	BaseEntity
}
