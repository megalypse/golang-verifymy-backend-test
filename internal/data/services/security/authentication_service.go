package security

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type AuthenticationService interface {
	GenerateJwtToken(userEmail string) (string, *models.CustomError)
}
