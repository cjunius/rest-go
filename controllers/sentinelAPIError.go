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
	ErrEntityNotFound     = &sentinelAPIError{status: http.StatusNotFound, msg: "entity not found"}
	ErrEntityTypeNotFound = &sentinelAPIError{status: http.StatusNotFound, msg: "entity type not found"}
)
