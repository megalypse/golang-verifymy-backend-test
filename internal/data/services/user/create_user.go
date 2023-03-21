package service

import (
	"crypto/rand"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

func (us UserService) Create(source *models.User) (*models.User, *models.CustomError) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on generating salt",
			Source:  nil,
		}
	}

	password := source.UserPassword.Password
	protectedPassword, err := bcrypt.GenerateFromPassword(append(password, salt...), bcrypt.DefaultCost)
	if err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on hashing password",
			Source:  nil,
		}
	}

	connection := us.userRepository.NewConnection()
	defer connection.CloseConnection()

	tx, cErr := connection.BeginTransaction()
	if cErr != nil {
		return nil, cErr
	}

	userId, cErr := us.userRepository.Create(tx, source)
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	userPassword := models.UserPassword{
		UserId:   userId,
		Password: protectedPassword,
		Salt:     salt,
	}

	userPasswordId, cErr := us.userPasswordRepository.Create(tx, &userPassword)
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	savedUserPassword, cErr := us.userPasswordRepository.FindById(tx, userPasswordId)
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	savedUser, cErr := us.userRepository.FindById(tx, userId)
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	savedUser.UserPassword = savedUserPassword

	tx.Commit()
	return savedUser, nil
}
