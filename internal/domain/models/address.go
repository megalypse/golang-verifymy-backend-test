package models

type Address struct {
	AddressAlias string `json:"alias" db:"address_alias"`
	ZipCode      string `json:"zipcode" db:"zip_code"`
	StreetName   string `json:"street_name" db:"street_name"`
	Number       string `json:"number" db:"number"`
	State        string `json:"state" db:"state"`
	Country      string `json:"country" db:"country"`

	UserId int64 `json:"user_id" db:"user_id"`

	BaseEntity
}
