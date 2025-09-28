package model

type Error interface {
	error
	ErrorCode() string
}

type AlreadExistsError struct {
	Code string
}

func (e *AlreadExistsError) Error() string {
	return e.Code
}

func (e *AlreadExistsError) ErrorCode() string {
	return e.Code
}
