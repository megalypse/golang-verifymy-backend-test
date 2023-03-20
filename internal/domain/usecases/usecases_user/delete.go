package usecases_user

type DeleteUser interface {
	Delete(int64) bool
}
