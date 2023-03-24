package auth

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

type EmailSignInService struct {
	userPasswordRepository repository.UserPasswordRepository
	userRepository         repository.UserRepository
	securityService        security.SecurityService
}

func NewUserEmailAuth(
	userPasswordRepository repository.UserPasswordRepository,
	userRepository repository.UserRepository,
	securityService security.SecurityService,
) *EmailSignInService {
	return &EmailSignInService{
		userPasswordRepository: userPasswordRepository,
		userRepository:         userRepository,
		securityService:        securityService,
	}
}

func (ua EmailSignInService) SignIn(ctx context.Context, source *models.User) (*models.User, *models.CustomError) {
	conn := factory.NewSqlConnection(ctx)
	defer conn.CloseConnection()

	tx, err := conn.BeginTransaction()
	if err != nil {
		return nil, err
	}

	user, err := ua.userRepository.FindByEmail(tx, source.Email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	storedPassword, err := ua.userPasswordRepository.FindLatestByUserId(tx, user.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	err = ua.securityService.Compare(storedPassword.Password, source.UserPassword.Password)
	if err != nil {
		tx.Commit()
		conn.CloseConnection()
		return nil, err
	}

	return user, nil
}
