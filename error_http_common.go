package errors

import "net/http"

func NewBadRequestError(error Error) HttpError {
	return NewHttpError(http.StatusBadRequest, error)
}

func NewInternalServerErrorError(error Error) HttpError {
	return NewHttpError(http.StatusInternalServerError, error)
}

func NewUnauthorizedError(error Error) HttpError {
	return NewHttpError(http.StatusUnauthorized, error)
}

func NewForbiddenError(error Error) HttpError {
	return NewHttpError(http.StatusForbidden, error)
}

func NewNotFoundError(error Error) HttpError {
	return NewHttpError(http.StatusNotFound, error)
}