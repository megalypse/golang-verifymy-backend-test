package security

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type SecurityService interface {
	SecureUserPassword(*models.UserPassword) (*models.UserPassword, *models.CustomError)
	Compare(hashedPassword []byte, plainpassword []byte) *models.CustomError
	// ValidateEmail(string) bool
}
