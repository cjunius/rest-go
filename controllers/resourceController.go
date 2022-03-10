package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PATH_PARAM_ENTITY_TYPE = "entityType"
const PATH_PARAM_ID = "id"

func CreateEntity(w http.ResponseWriter, r *http.Request) {
	entityType := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]

	decoder := json.NewDecoder(r.Body)
	var entity map[string]interface{}
	if err := decoder.Decode(&entity); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if savedEntity, err := AddEntityData(entityType, entity); err != nil {
		statusCode, msg := err.APIError()
		http.Error(w, msg, statusCode)
		return
	} else {

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(savedEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func GetEntities(w http.ResponseWriter, r *http.Request) {
	entityType := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]

	if entities, err := GetEntitiesData(entityType); err != nil {
		statusCode, msg := err.APIError()
		http.Error(w, msg, statusCode)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(entities); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func GetEntity(w http.ResponseWriter, r *http.Request) {
	entityType := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]

	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {

		if entity, err := GetEntityData(entityType, id); err != nil {
			statusCode, msg := err.APIError()
			http.Error(w, msg, statusCode)
			return
		} else {
			if err := json.NewEncoder(w).Encode(entity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func ReplaceEntity(w http.ResponseWriter, r *http.Request) {
	entityType := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {

		decoder := json.NewDecoder(r.Body)
		var entity map[string]interface{}
		if err := decoder.Decode(&entity); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if replacedEntity, err := ReplaceEntityData(entityType, id, entity); err != nil {
			statusCode, msg := err.APIError()
			http.Error(w, msg, statusCode)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(replacedEntity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
}

func UpdateEntity(w http.ResponseWriter, r *http.Request) {
	entityType := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {

		decoder := json.NewDecoder(r.Body)
		var entity map[string]interface{}
		if err := decoder.Decode(&entity); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if updatedEntity, err := UpdateEntityData(entityType, id, entity); err != nil {
			statusCode, msg := err.APIError()
			http.Error(w, msg, statusCode)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(updatedEntity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}

}

func DeleteEntity(w http.ResponseWriter, r *http.Request) {
	entityType := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {
		if err := RemoveEntityData(entityType, id); err != nil {
			statusCode, msg := err.APIError()
			http.Error(w, msg, statusCode)
			return
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
