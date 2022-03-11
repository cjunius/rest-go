package controllers

import (
	"testing"
)

var entity = map[string]interface{}{
	"name": "Test Entity",
}

var entity2 = map[string]interface{}{
	"name": "Test Entity 2",
}

var entity3 = map[string]interface{}{
	"newattribute": "newvalue",
}

func TestAddEntitySuccess(t *testing.T) {
	if actual, err := AddEntityData("Persons", entity); err != nil {
		t.Errorf("No Error Expected but received error %s", err)
	} else {
		if actual["name"] != "Test Entity" {
			t.Errorf("Didn't save and retrieve data correctly")
		}

		if actual["id"] == nil {
			t.Errorf("Didn't generate id value correctly")
		}

		t.Cleanup(func() {
			id := actual["id"].(int)
			RemoveEntityData("Persons", id)
		})
	}
}

func TestGetEntities_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	if _, err := GetEntitiesData(entityType); err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}
}

func TestGetEntities_Success(t *testing.T) {
	entityType := "Persons"
	person1, _ := AddEntityData(entityType, entity)
	person2, _ := AddEntityData(entityType, entity)

	if entities, err := GetEntitiesData(entityType); err != nil {
		_, msg := err.APIError()
		t.Errorf("Did not expect error, but received %s", msg)
	} else {
		if len(entities) != 2 {
			t.Errorf("Expected 2 entities, but received %d", len(entities))
		}
	}

	t.Cleanup(func() {
		id1 := person1["id"].(int)
		id2 := person2["id"].(int)
		RemoveEntityData("Persons", id1)
		RemoveEntityData("Persons", id2)
	})
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
	returnedEntity, _ := AddEntityData(entityType, entity)

	if _, err := GetEntityData(entityType, 20); err != ErrEntityNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Not Found but received %s", msg)
	}

	t.Cleanup(func() {
		id := returnedEntity["id"].(int)
		RemoveEntityData(entityType, id)
	})
}

func TestGetEntity_Success(t *testing.T) {
	entityType := "Exists"
	addedEntity, _ := AddEntityData(entityType, entity)
	id := addedEntity["id"].(int)

	if returnedEntity, err := GetEntityData(entityType, id); err != nil {
		_, msg := err.APIError()
		t.Errorf("Did not expecte error but received %s", msg)
	} else {

		if returnedEntity["id"].(int) != id {
			t.Errorf("Did not successfully retrieve entity")
		}

		t.Cleanup(func() {
			id := returnedEntity["id"].(int)
			RemoveEntityData(entityType, id)
		})

	}
}

func TestReplaceEntitiy_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	_, err := ReplaceEntityData(entityType, 1, entity)

	if err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}
}

func TestReplaceEntitiy_EntityNotFound(t *testing.T) {
	entityType := "Exists"
	returnedEntity, _ := AddEntityData(entityType, entity)

	_, err := ReplaceEntityData(entityType, 20, entity)

	if err != ErrEntityNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Not Found but received %s", msg)
	}

	t.Cleanup(func() {
		id := returnedEntity["id"].(int)
		RemoveEntityData(entityType, id)
	})
}

func TestReplaceEntitiy_Success(t *testing.T) {
	entityType := "Persons"
	addedEntity, _ := AddEntityData(entityType, entity)

	returnedEntity, _ := ReplaceEntityData(entityType, addedEntity["id"].(int), entity2)

	if returnedEntity["name"] != entity2["name"] {
		t.Errorf("Entity was not successfully replaced ")
	}

	t.Cleanup(func() {
		id := returnedEntity["id"].(int)
		RemoveEntityData(entityType, id)
	})
}

func TestUpdateEntity_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	_, err := UpdateEntityData(entityType, 1, entity)

	if err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}
}

func TestUpdateEntity_EntityNotFound(t *testing.T) {
	entityType := "Exists"
	returnedEntity, _ := AddEntityData(entityType, entity)

	_, err := UpdateEntityData(entityType, 20, entity)

	if err != ErrEntityNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Not Found but received %s", msg)
	}

	t.Cleanup(func() {
		id := returnedEntity["id"].(int)
		RemoveEntityData(entityType, id)
	})
}

func TestUpdateEntitiy_Success(t *testing.T) {
	entityType := "Persons"
	addedEntity, _ := AddEntityData(entityType, entity)

	returnedEntity, _ := UpdateEntityData(entityType, addedEntity["id"].(int), entity3)

	if returnedEntity["name"] != entity["name"] {
		t.Errorf("Entity changed attribute value is should not have")
	}

	if returnedEntity["newattribute"] != entity3["newattribute"] {
		t.Errorf("Entity was NOT successfully updated")
	}

	t.Cleanup(func() {
		id := returnedEntity["id"].(int)
		RemoveEntityData(entityType, id)
	})
}

func TestRemoveEntity_EntityTypeNotFound(t *testing.T) {
	entityType := "NotFound"

	if err := RemoveEntityData(entityType, 1); err != ErrEntityTypeNotFound {
		_, msg := err.APIError()
		t.Errorf("Expected Error Entity Type Not Found but received %s", msg)
	}
}

func TestRemoveEntity_Success(t *testing.T) {
	entityType := "Exists"
	returnedEntity, _ := AddEntityData(entityType, entity)

	id := returnedEntity["id"].(int)
	if err := RemoveEntityData(entityType, id); err != nil {
		_, msg := err.APIError()
		t.Errorf("Did not expect error but got one, %s", msg)
	}
}
