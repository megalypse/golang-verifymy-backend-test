package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/db/repositorymysql"
)

var userPasswordRepository repository.UserPasswordRepository

func init() {
	userPasswordRepository = repositorymysql.MySqlUserPasswordRepository{}
}

func GetMySqlUserPasswordRepository() repository.UserPasswordRepository {
	return userPasswordRepository
}
