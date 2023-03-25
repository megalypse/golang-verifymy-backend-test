package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	repositorymysql "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/roles"
)

var rolesRepository repository.RolesRepository

func init() {
	rolesRepository = repositorymysql.MySqlRolesRepository{}
}

func GetRolesRepository() repository.RolesRepository {
	return rolesRepository
}
