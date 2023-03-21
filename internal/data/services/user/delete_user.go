package service

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

func (us UserService) Delete(id int64) *models.CustomError {
	connection := us.userRepository.NewConnection()
	defer connection.CloseConnection()

	tx, err := connection.BeginTransaction()
	if err != nil {
		return err
	}

	if err := us.userRepository.Delete(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
