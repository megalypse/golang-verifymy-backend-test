package service

import (
	"context"
	"crypto/rand"
	"log"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

func (us UserService) Create(ctx context.Context, source *models.User) (*models.User, *models.CustomError) {
	protectedUserPassword, err := secureUserPassword(source)
	if err != nil {
		return nil, err
	}

	connection := us.userRepository.NewConnection(ctx)

	writeTx, err := connection.BeginTransaction()
	if err != nil {
		return nil, err
	}

	userId, err := us.userRepository.Create(writeTx, source)
	if err != nil {
		writeTx.Rollback()
		return nil, err
	}

	_, err = us.saveUserPassword(writeTx, models.UserPassword{
		UserId:   userId,
		Password: protectedUserPassword.Password,
		Salt:     protectedUserPassword.Salt,
	})
	if err != nil {
		writeTx.Rollback()
		return nil, err
	}

	err = us.saveUserAddressList(writeTx, source.AddressList, userId)
	if err != nil {
		writeTx.Rollback()
		return nil, err
	}

	err = writeTx.Commit()
	if err != nil {
		writeTx.Rollback()
		return nil, err
	}

	readTx, err := connection.BeginTransaction()
	if err != nil {
		return nil, err
	}

	fullSavedUser, err := us.getUserWithAddresses(readTx, userId)
	if err != nil {
		readTx.Rollback()
		return nil, err
	}

	if err = readTx.Commit(); err != nil {
		return nil, err
	}

	if err = connection.CloseConnection(); err != nil {
		return nil, err
	}

	log.Println("User created")
	return fullSavedUser, nil
}

func secureUserPassword(source *models.User) (*models.UserPassword, *models.CustomError) {
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

	return &models.UserPassword{
		Password: protectedPassword,
		Salt:     salt,
	}, nil
}

func (us UserService) getUserWithAddresses(tx repository.Transaction, userId int64) (*models.User, *models.CustomError) {
	savedAddressList, cErr := us.addressRepository.GetAllByUserId(tx, userId)
	if cErr != nil {
		return nil, cErr
	}

	savedUser, cErr := us.userRepository.FindById(tx, userId)
	if cErr != nil {
		return nil, cErr
	}

	savedUser.AddressList = savedAddressList

	return savedUser, nil
}

func (us UserService) saveUserAddressList(tx repository.Transaction, source []models.Address, userId int64) *models.CustomError {
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
