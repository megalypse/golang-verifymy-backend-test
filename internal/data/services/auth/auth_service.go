package auth

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UserEmailAuth struct {
	userPasswordRepository repository.UserPasswordRepository
	userRepository         repository.UserRepository
	securityService        security.SecurityService
}

func NewUserEmailAuth(
	userPasswordRepository repository.UserPasswordRepository,
	userRepository repository.UserRepository,
	securityService security.SecurityService,
) *UserEmailAuth {
	return &UserEmailAuth{
		userPasswordRepository: userPasswordRepository,
		userRepository:         userRepository,
		securityService:        securityService,
	}
}

func (ua UserEmailAuth) Auth(ctx context.Context, source *models.User) (bool, *models.CustomError) {
	conn := ua.userRepository.NewConnection(ctx)
	defer conn.CloseConnection()

	tx, err := conn.BeginTransaction()
	if err != nil {
		return false, err
	}

	user, err := ua.userRepository.FindByEmail(tx, source.Email)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	storedPassword, err := ua.userPasswordRepository.FindLatestByUserId(tx, user.Id)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()

	err = ua.securityService.Compare(storedPassword.Password, source.UserPassword.Password)
	if err != nil {
		tx.Commit()
		conn.CloseConnection()
		return false, err
	}

	return true, nil
}
