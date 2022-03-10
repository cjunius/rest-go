package controllers

/* global variables */
var data = make(map[string][]map[string]interface{})
var nextId = 1

func AddEntityData(entityType string, entity map[string]interface{}) (map[string]interface{}, APIError) {
	entity["id"] = nextId
	if _, ok := data[entityType]; !ok {
		data[entityType] = make([]map[string]interface{}, 0)
	}
	data[entityType] = append(data[entityType], entity)
	id := nextId
	nextId++
	return GetEntityData(entityType, id)
}

func GetEntitiesData(entityType string) ([]map[string]interface{}, APIError) {
	if val, ok := data[entityType]; ok {
		return val, nil
	}
	return nil, ErrEntityTypeNotFound
}

func GetEntityData(entityType string, id int) (map[string]interface{}, APIError) {

	if entities, err := GetEntitiesData(entityType); err != nil {
		return nil, err
	} else {
		for i := range entities {
			if entities[i]["id"] == id {
				return entities[i], nil
			}
		}
		return nil, ErrEntityNotFound
	}
}

func UpdateEntityData(entityType string, id int, entity map[string]interface{}) (map[string]interface{}, APIError) {
	if entities, err := GetEntitiesData(entityType); err != nil {
		return nil, err
	} else {

		for i := range entities {
			if entities[i]["id"] == id {

				for k, _ := range entity {
					entities[i][k] = entity[k]
					entities[i]["id"] = id
				}

				return entities[i], nil
			}
		}
		return nil, ErrEntityNotFound
	}
}

func RemoveEntityData(entityType string, id int) APIError {
	if entities, err := GetEntitiesData(entityType); err != nil {
		return err
	} else {
		if len(entities) != 0 {
			index := 0
			for i := range entities {
				if entities[i]["id"] == id {
					index = i
				}
			}
			copy(entities[index:], entities[index+1:])
			entities[len(entities)-1] = nil
			data[entityType] = entities[:len(entities)-1]
		}
	}
	return nil
}
