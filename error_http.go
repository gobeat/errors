package errors

type HttpError interface {
	Error() Error
	WithError(error Error) HttpError
	Status() int
	WithStatus(status int) HttpError
}

func NewHttpError(status int, error Error) HttpError {
	return &factoryHttpError{
		status: status,
		error:  error,
	}
}

type factoryHttpError struct {
	status int
	error  Error
}

func (e *factoryHttpError) Error() Error {
	return e.error
}

func (e *factoryHttpError) WithError(error Error) HttpError {
	e.error = error
	return e
}

func (e *factoryHttpError) Status() int {
	return e.status
}

func (e *factoryHttpError) WithStatus(status int) HttpError {
	e.status = status
	return e
}
