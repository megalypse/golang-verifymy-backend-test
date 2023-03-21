package models

type Address struct {
	ZipCode    string
	StreetName string
	Number     int64
	State      string
	Country    string

	BaseEntity
}
