package models

type UserPassword struct {
	Password []byte `json:"password"`
	Salt     []byte

	UserId int64

	Identity
	CreatedDate
}
