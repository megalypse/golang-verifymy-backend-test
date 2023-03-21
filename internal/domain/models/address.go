package models

type Address struct {
	AddressAlias string `json:"alias"`
	ZipCode      string `json:"zipcode"`
	StreetName   string `json:"street_name"`
	Number       string `json:"number"`
	State        string `json:"state"`
	Country      string `json:"country"`

	UserId int64 `json:"user_id"`

	BaseEntity
}
