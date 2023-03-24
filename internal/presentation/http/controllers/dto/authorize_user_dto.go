package dto

type AuthorizeUserDto struct {
	UserId int64 `json:"user_id" example:"1"`
	RoleId int64 `json:"role_id" example:"1"`
}
