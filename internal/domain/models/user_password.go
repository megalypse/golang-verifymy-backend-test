package models

type UserPassword struct {
	Password []byte
	Salt     []byte

	UserId int64

	Identity
	CreatedAt
}
