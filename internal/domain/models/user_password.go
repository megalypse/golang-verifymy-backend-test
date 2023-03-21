package models

type UserPassword struct {
	UserId   int64
	Password []byte
	Salt     []byte

	Identity
	CreatedAt
}
