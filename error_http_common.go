package errors

import "net/http"

func NewBadRequestError(code string, message string) HttpError {
	return NewHttpError(code, message, http.StatusBadRequest)
}

func NewInternalServerErrorError(code string, message string) HttpError {
	return NewHttpError(code, message, http.StatusInternalServerError)
}

func NewUnauthorizedError(code string, message string) HttpError {
	return NewHttpError(code, message, http.StatusUnauthorized)
}

func NewForbiddenError(code string, message string) HttpError {
	return NewHttpError(code, message, http.StatusForbidden)
}

func NewNotFoundError(code string, message string) HttpError {
	return NewHttpError(code, message, http.StatusNotFound)
}
