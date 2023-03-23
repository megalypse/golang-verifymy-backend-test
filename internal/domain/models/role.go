package models

type Role struct {
	Identity
	CreatedDate

	Alias string `json:"alias"`
}
