package controllers

type APIError interface {
	APIError() (int, string)
}
