package errors

type HttpError interface {
	Error() Error
	WithError(error Error) HttpError
	Status() int
	WithStatus(status int) HttpError
}

func NewHttpError(status int, error Error) HttpError {
	return &FactoryHttpError{
		status: status,
		error:  error,
	}
}

type FactoryHttpError struct {
	status int
	error  Error
}

func (e *FactoryHttpError) Error() Error {
	return e.error
}

func (e *FactoryHttpError) WithError(error Error) HttpError {
	e.error = error
	return e
}

func (e *FactoryHttpError) Status() int {
	return e.status
}

func (e *FactoryHttpError) WithStatus(status int) HttpError {
	e.status = status
	return e
}
