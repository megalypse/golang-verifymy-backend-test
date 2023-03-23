package models

type UserPassword struct {
	Password []byte `json:"password"`

	UserId int64

	Identity
	CreatedDate
}
