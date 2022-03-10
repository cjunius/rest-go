package controllers

import "net/http"

type sentinelAPIError struct {
	status int
	msg    string
}

func (e sentinelAPIError) Error() string {
	return e.msg
}

func (e sentinelAPIError) Status() int {
	return e.status
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.msg
}

var (
	ErrAuth               = &sentinelAPIError{status: http.StatusUnauthorized, msg: "invalid token"}
	ErrEntityNotFound     = &sentinelAPIError{status: http.StatusNotFound, msg: "entity not found"}
	ErrEntityTypeNotFound = &sentinelAPIError{status: http.StatusNotFound, msg: "entity type not found"}
	ErrDuplicate          = &sentinelAPIError{status: http.StatusBadRequest, msg: "duplicate"}
	ErrInvalidJSON        = &sentinelAPIError{status: http.StatusBadRequest, msg: "invalid json"}
	ErrInvalidId          = &sentinelAPIError{status: http.StatusBadRequest, msg: "invalid id"}
)
