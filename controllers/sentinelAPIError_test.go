package controllers

import (
	"net/http"
	"testing"
)

func TestErrAuth(t *testing.T) {
	if ErrAuth.Error() != "invalid token" {
		t.Errorf("Enum error message is incorrect")
	}

	if ErrAuth.Status() != http.StatusUnauthorized {
		t.Errorf("Enum status is incorrect")
	}

	status, msg := ErrAuth.APIError()
	if status != http.StatusUnauthorized {
		t.Errorf("Enum status is incorrect")
	}

	if msg != "invalid token" {
		t.Errorf("Enum error message is incorrect")
	}
}
