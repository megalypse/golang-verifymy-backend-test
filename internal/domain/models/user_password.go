package models

type UserPassword struct {
	Password []byte `json:"password" db:"password" swaggertype:"string" example:"johnspassword+123"`

	UserId int64 `json:"user_id" db:"user_id" swaggerignore:"true"`

	Identity
	CreatedDate
}
