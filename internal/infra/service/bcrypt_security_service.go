package service

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type BCryptSecurityService struct{}

func (BCryptSecurityService) SecureUserPassword(source *models.UserPassword) (*models.UserPassword, *models.CustomError) {
	protectedPassword, err := hashPassword(source.Password)
	if err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on hashing password",
			Source:  nil,
		}
	}

	return &models.UserPassword{
		Password: protectedPassword,
	}, nil
}

func (BCryptSecurityService) Compare(hashedPassword []byte, plainpassword []byte) *models.CustomError {
	err := bcrypt.CompareHashAndPassword(hashedPassword, plainpassword)
	if err != nil {
		return &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	return nil
}

func hashPassword(plainPassword []byte) ([]byte, *models.CustomError) {
	hashedPassword, err := bcrypt.GenerateFromPassword(plainPassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return hashedPassword, nil
}
