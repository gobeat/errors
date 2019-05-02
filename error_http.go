package errors

type HttpError interface {
	Error
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

func (e *factoryHttpError) Code() string {
	return e.error.Code()
}

func (e *factoryHttpError) WithCode(code string) Error {
	e.error.WithCode(code)
	return e
}

func (e *factoryHttpError) Message() string {
	return e.error.Message()
}

func (e *factoryHttpError) WithMessage(message string) Error {
	e.error.WithMessage(message)
	return e
}

func (e *factoryHttpError) Status() int {
	return e.status
}

func (e *factoryHttpError) WithStatus(status int) HttpError {
	e.status = status
	return e
}
