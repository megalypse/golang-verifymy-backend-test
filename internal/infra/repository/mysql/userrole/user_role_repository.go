package repository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

type MySqlUserRolesRepository struct {
}

func (MySqlUserRolesRepository) GetAllByUserId(tx repository.Transaction, userid int64) ([]models.Role, *models.CustomError) {
	result, err := tx.Query("SELECT roles.id, roles.alias, roles.created_at FROM user_roles JOIN roles ON user_roles.role_id = roles.id WHERE user_roles.user_id = ?", userid)

	if err != nil {
		return nil, err
	}

	userRoles, err := mappers.MapMany(mappers.RoleMapperFunc, result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	return userRoles, nil
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
