package service

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type RolesAuthorizationService struct {
	RolesRepository     repository.RolesRepository
	UserRolesRepository repository.UserRolesRepository
}

func (as RolesAuthorizationService) GetUserRoles(ctx context.Context, userId int64) ([]models.Role, *models.CustomError) {
	conn := as.RolesRepository.NewConnection(ctx)
	defer conn.CloseConnection()

	tx, err := conn.BeginTransaction()
	if err != nil {
		return nil, err
	}

	roles, err := as.UserRolesRepository.GetAllByUserId(tx, userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return roles, nil
}

func (as RolesAuthorizationService) AssignRole(ctx context.Context, userId int64, roleAlias string) *models.CustomError {
	conn := as.RolesRepository.NewConnection(ctx)
	defer conn.CloseConnection()

	tx, err := conn.BeginTransaction()
	if err != nil {
		return err
	}

	role, err := as.RolesRepository.FindByAlias(tx, roleAlias)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = as.UserRolesRepository.AssignRole(tx, userId, role.Id); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
