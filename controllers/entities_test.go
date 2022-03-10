package controllers

import (
	"testing"
)

var entity = map[string]interface{}{
	"name": "Test Entity",
}

func TestAddEntitySuccess(t *testing.T) {
	if actual, err := AddEntityData("Persons", entity); err != nil {
		t.Errorf("No Error Expected but recieved error %s", err)
	} else {
		if actual["name"] != "Test Entity" {
			t.Errorf("Didn't save and retrieve data correctly")
		}

		if actual["id"] == nil {
			t.Errorf("Didn't generate id value correctly")
		}
	}
}

func TestGetEntities_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	if _, err := GetEntitiesData(entityType); err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}

}

func TestGetEntity_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	if _, err := GetEntityData(entityType, 1); err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}

}

func TestGetEntity_EntityNotFound(t *testing.T) {
	entityType := "Exists"
	AddEntityData(entityType, entity)

	if _, err := GetEntityData(entityType, 20); err != ErrEntityNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Not Found but received %s", msg)
	}

}

func TestUpdateEntity_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	_, err := UpdateEntityData(entityType, 1, entity)

	if err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}

}

func TestRemoveEntity_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	if err := RemoveEntityData(entityType, 1); err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}

}
