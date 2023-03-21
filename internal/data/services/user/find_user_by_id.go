package service

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

func (us UserService) FindById(id int64) (*models.User, *models.CustomError) {
	connection := us.userRepository.NewConnection()
	defer connection.CloseConnection()

	tx, err := connection.BeginTransaction()
	if err != nil {
		return nil, err
	}

	result, err := us.userRepository.FindById(tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return result, nil
}
