package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	mySqlUserRepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/user"
)

var userRepository repository.UserRepository

func init() {
	userRepository = mySqlUserRepository.MySqlUserRepository{}
}

func GetUserRepository() repository.UserRepository {
	return userRepository
}
