package service

import (
	"context"
	"crypto/rand"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

func (us UserService) Create(ctx context.Context, source *models.User) (*models.User, *models.CustomError) {
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

	connection := us.userRepository.NewConnection(ctx)
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

	_, cErr = us.saveUserPassword(tx, models.UserPassword{
		UserId:   userId,
		Password: protectedPassword,
		Salt:     salt,
	})
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	cErr = us.saveAddresses(tx, source.AddressList, userId)
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	cErr = tx.Commit()
	if cErr != nil {
		tx.Rollback()
		return nil, cErr
	}

	rtx, cErr := connection.BeginTransaction()
	if cErr != nil {
		return nil, cErr
	}

	fullSavedUser, cErr := us.getFullSavedUser(rtx, userId)
	if cErr != nil {
		rtx.Rollback()
		return nil, cErr
	}

	tx.Commit()
	return fullSavedUser, nil
}

func (us UserService) getFullSavedUser(tx repository.Transaction, userId int64) (*models.User, *models.CustomError) {
	savedUserPassword, cErr := us.userPasswordRepository.FindLatestByUserId(tx, userId)
	if cErr != nil {
		return nil, cErr
	}

	savedAddressList, cErr := us.addressRepository.GetAllByUserId(tx, userId)
	if cErr != nil {
		return nil, cErr
	}

	savedUser, cErr := us.userRepository.FindById(tx, userId)
	if cErr != nil {
		return nil, cErr
	}

	savedUser.UserPassword = savedUserPassword
	savedUser.AddressList = savedAddressList

	return savedUser, nil
}

func (us UserService) saveAddresses(tx repository.Transaction, source []models.Address, userId int64) *models.CustomError {
	addressList := make([]models.Address, 0, len(source))
	for _, address := range source {
		address.UserId = userId

		addressList = append(addressList, address)
	}

	for _, address := range addressList {
		_, cErr := us.addressRepository.Create(tx, &address)
		if cErr != nil {
			return cErr
		}
	}

	return nil
}

func (us UserService) saveUserPassword(tx repository.Transaction, source models.UserPassword) (int64, *models.CustomError) {
	userPasswordId, cErr := us.userPasswordRepository.Create(tx, &source)
	if cErr != nil {
		return 0, cErr
	}

	return userPasswordId, nil
}
