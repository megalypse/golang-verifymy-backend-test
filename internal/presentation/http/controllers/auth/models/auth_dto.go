package models

type AuthDto struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
