package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	repositorymysql "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
)

var userPasswordRepository repository.UserPasswordRepository

func init() {
	userPasswordRepository = repositorymysql.MySqlUserPasswordRepository{}
}

func GetMySqlUserPasswordRepository() repository.UserPasswordRepository {
	return userPasswordRepository
}
