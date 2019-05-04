package errors

type HttpError interface {
	Error
	Status() int
	WithStatus(status int) HttpError
}

func NewHttpError(code string, message string, status int) HttpError {
	return &factoryHttpError{
		status: status,
		factoryError: factoryError{
			code:    code,
			message: message,
		},
	}
}

type factoryHttpError struct {
	factoryError
	status int
}

func (e *factoryHttpError) Status() int {
	return e.status
}

func (e *factoryHttpError) WithStatus(status int) HttpError {
	e.status = status
	return e
}
