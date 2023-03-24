package repository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

type MySqlRolesRepository struct {
}

func (rp MySqlRolesRepository) FindById(tx repository.Transaction, roleId int64) (*models.Role, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM roles
	WHERE id = ?
	LIMIT 1
	`, roleId)

	if err != nil {
		return nil, err
	}

	role, err := mappers.MapOne(mappers.RoleMapperFunc, result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	return role, nil
}
