package services

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (us UserService) Create(source *models.User) (*models.User, *models.CustomError) {
	return us.userRepository.Create(source)
}

func (us UserService) Delete(id int64) *models.CustomError {
	return us.userRepository.Delete(id)
}

func (us UserService) FindById(id int64) (*models.User, *models.CustomError) {
	return us.userRepository.FindById(id)
}

func (us UserService) Update(source *models.User) (*models.User, *models.CustomError) {
	return us.userRepository.Update(source)
}
