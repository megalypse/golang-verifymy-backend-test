package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/db/repositorymysql"
)

var userRepository repository.UserRepository

func init() {
	userRepository = repositorymysql.MySqlUserRepository{}
}

func GetUserRepository() repository.UserRepository {
	return userRepository
}
