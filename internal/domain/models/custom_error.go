package models

type CustomError struct {
	Code    int
	Message string
	Source  error
}

func (ce CustomError) Error() string {
	return ce.Message
}
