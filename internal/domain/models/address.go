package models

type Address struct {
	AddressAlias string `db:"address_alias" json:"address_alias"`
	ZipCode      string `db:"zip_code" json:"zip_code"`
	StreetName   string `db:"street_name" json:"street_name"`
	Number       string `db:"number" json:"number"`
	State        string `db:"state" json:"state"`
	Country      string `db:"country" json:"country"`

	UserId int64 `json:"user_id" db:"user_id"`

	BaseEntity
}
