package service

import (
	"crypto/rand"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type BCryptSecurityService struct{}

func (BCryptSecurityService) SecureUserPassword(source models.UserPassword) (*models.UserPassword, *models.CustomError) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on generating salt",
			Source:  nil,
		}
	}
	password := source.Password
	protectedPassword, err := bcrypt.GenerateFromPassword(append(password, salt...), bcrypt.DefaultCost)
	if err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on hashing password",
			Source:  nil,
		}
	}

	return &models.UserPassword{
		Password: protectedPassword,
		Salt:     salt,
	}, nil
}
