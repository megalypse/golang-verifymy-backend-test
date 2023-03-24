package models

type Address struct {
	AddressAlias string `db:"address_alias"`
	ZipCode      string `db:"zip_code"`
	StreetName   string `db:"street_name"`
	Number       string `db:"number"`
	State        string `db:"state"`
	Country      string `db:"country"`

	UserId int64 `json:"user_id" db:"user_id" example:"1"`

	BaseEntity
}
