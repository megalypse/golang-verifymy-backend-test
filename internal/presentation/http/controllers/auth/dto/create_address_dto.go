package dto

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type CreateAddressDto struct {
	Alias      string `json:"alias" example:"Home"`
	ZipCode    string `json:"zipcode" example:"0000000"`
	StreetName string `json:"street_name" example:"Generic Street Name"`
	Number     string `json:"number" example:"GN-01"`
	State      string `json:"state" example:"GT"`
	Country    string `json:"country" example:"GC"`
	UserId     int64  `json:"user_id" example:"1"`
}

func (ad CreateAddressDto) ToAddress() *models.Address {
	return &models.Address{
		AddressAlias: ad.Alias,
		ZipCode:      ad.ZipCode,
		StreetName:   ad.StreetName,
		Number:       ad.Number,
		State:        ad.State,
		Country:      ad.Country,
		UserId:       0,
	}
}
