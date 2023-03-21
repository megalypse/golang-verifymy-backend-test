package models

type Address struct {
	AddressAlias string
	ZipCode      string
	StreetName   string
	Number       int64
	State        string
	Country      string

	UserId int64

	BaseEntity
}
