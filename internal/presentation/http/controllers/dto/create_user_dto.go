package dto

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type CreateUserDto struct {
	Name     string           `json:"name" example:"John Doe"`
	Age      int              `json:"age" example:"36"`
	Email    string           `json:"email" example:"johndoe36@genericmail.com"`
	Password string           `json:"password" example:"genericPassword+123"`
	Address  CreateAddressDto `json:"address"`
}

func (ud CreateUserDto) ToUserModel() *models.User {
	return &models.User{
		Name:  ud.Name,
		Age:   ud.Age,
		Email: ud.Email,
		UserPassword: &models.UserPassword{
			Password: []byte(ud.Password),
		},
		AddressList: []models.Address{
			*ud.Address.ToAddress(),
		},
	}
}
