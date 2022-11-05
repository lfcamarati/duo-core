package handler

import "net/http"

type ResponseError interface {
	Code() int
	Message() string
}

type JsonResponseError struct {
	httpCode int
	Error    string `json:"message"`
}

func (e JsonResponseError) Code() int {
	return e.httpCode
}

func (e JsonResponseError) Message() string {
	return e.Error
}

func NewJsonResponseError(code int, error string) ResponseError {
	return JsonResponseError{code, error}
}

func NewNotAuthorizedError(error string) ResponseError {
	return NewJsonResponseError(http.StatusUnauthorized, error)
}

func NewUsecaseError(error string) ResponseError {
	return NewJsonResponseError(http.StatusUnprocessableEntity, error)
}

func NewBadRequest(error string) ResponseError {
	return NewJsonResponseError(http.StatusBadRequest, error)
}

func NewInternalServerError(error string) ResponseError {
	return NewJsonResponseError(http.StatusInternalServerError, error)
}

func NewNotFoundError() ResponseError {
	return NewJsonResponseError(http.StatusNotFound, "Registro não encontrado")
}
