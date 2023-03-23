package repository

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MySqlUserRolesRepository struct {
}

func (ur MySqlUserRolesRepository) AssignRole(tx repository.Transaction, userId, roleId int64) *models.CustomError {
	_, err := tx.Exec(`
	INSERT INTO users_roles (user_id, role_id)
	VALUES (?, ?);
	`, userId, roleId)

	if err != nil {
		return err
	}

	return nil
}
