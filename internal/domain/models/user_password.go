package models

type UserPassword struct {
	Password []byte `json:"password" db:"password"`

	UserId int64 `json:"user_id" db:"user_id"`

	Identity
	CreatedDate
}
