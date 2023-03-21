package models

type Security struct {
	UserId   int64
	Password string
	Salt     string

	Identity
	CreatedAt
	DeletedAt
}
