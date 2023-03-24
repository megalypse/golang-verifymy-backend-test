package dto

type AuthDto struct {
	Email    string `json:"email" example:"johndoe36@genericmail.com"`
	Password string `json:"password" example:"genericPassword+123"`
}
