package models

type User struct {
	Name         string        `json:"name"`
	Age          int           `json:"age"`
	Email        string        `json:"email"`
	UserPassword *UserPassword `json:"security"`
	AddressList  []Address     `json:"address_list"`

	BaseEntity
}
