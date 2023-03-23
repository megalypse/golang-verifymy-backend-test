package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	userrolesRepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/userrole"
)

var userRolesRepository repository.UserRolesRepository

func init() {
	userRolesRepository = userrolesRepository.MySqlUserRolesRepository{}
}

func GetMySqlUserRolesRepository() repository.UserRolesRepository {
	return userRolesRepository
}
