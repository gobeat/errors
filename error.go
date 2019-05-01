package errors

type Error interface {
	Code() string
	WithCode(code string) Error
	Message() string
	WithMessage(message string) Error
}

func NewError(code string, message string) Error {
	return &FactoryError{
		code:    code,
		message: message,
	}
}

type FactoryError struct {
	code    string
	message string
}

func (e *FactoryError) Code() string {
	return e.code
}

func (e *FactoryError) WithCode(code string) Error {
	e.code = code
	return e
}

func (e *FactoryError) Message() string {
	return e.message
}

func (e*FactoryError) WithMessage(message string) Error {
	e.message = message
	return e
}


