package repository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	repositorymysql "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

type MySqlRolesRepository struct {
	repositorymysql.BaseMySqlRepository
}

func (rp MySqlRolesRepository) FindByAlias(tx repository.Transaction, roleAlias string) (*models.Role, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM roles
	WHERE alias = ?
	LIMIT 1
	`, roleAlias)

	if err != nil {
		return nil, err
	}

	role, err := mappers.MapOne(mappers.RoleMapperFunc, result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	return role, nil
}
