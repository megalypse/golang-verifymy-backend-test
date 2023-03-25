package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	repositorymysql "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/userpassword"
)

var userPasswordRepository repository.UserPasswordRepository

func init() {
	userPasswordRepository = repositorymysql.MySqlUserPasswordRepository{}
}

func GetUserPasswordRepository() repository.UserPasswordRepository {
	return userPasswordRepository
}
