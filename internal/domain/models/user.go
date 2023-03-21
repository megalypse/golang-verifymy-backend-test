package models

type User struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Security Security `json:"security"`

	BaseEntity
}
