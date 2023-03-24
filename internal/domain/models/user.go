package models

type User struct {
	Name         string        `db:"name"`
	Age          int           `db:"age"`
	Email        string        `db:"email"`
	UserPassword *UserPassword `db:"user_password"`
	AddressList  []Address     `json:"address_list"`

	BaseEntity
}
