package models

type UserRole struct {
	UserId int64 `db:"user_id"`
	RoleId int64 `db:"role_id"`
}
