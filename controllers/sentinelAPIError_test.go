package controllers

import (
	"net/http"
	"testing"
)

func TestErrAuth(t *testing.T) {
	if ErrEntityNotFound.Error() != "entity not found" {
		t.Errorf("Enum error message is incorrect")
	}

	if ErrEntityNotFound.Status() != http.StatusNotFound {
		t.Errorf("Enum status is incorrect")
	}

	status, msg := ErrEntityNotFound.APIError()
	if status != http.StatusNotFound {
		t.Errorf("Enum status is incorrect")
	}

	if msg != "entity not found" {
		t.Errorf("Enum error message is incorrect")
	}
}
