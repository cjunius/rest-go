package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* global variables */
var data = make(map[string]map[int]map[string]interface{})
var nextId = 1

const PATH_PARAM_ENTITY_TYPE = "entityType"
const PATH_PARAM_ID = "id"

func CreateEntity(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]

	decoder := json.NewDecoder(r.Body)
	var t map[string]interface{}
	if err := decoder.Decode(&t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	t["id"] = nextId
	log.Println("Adding Entity")
	log.Println(t)

	if resourcesMap, ok := data[entity]; !ok {
		data[entity] = make(map[int]map[string]interface{})
		data[entity][nextId] = t
	} else {
		resourcesMap[nextId] = t
	}

	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nextId = nextId + 1
}

func GetEntities(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if entitiesMap, ok := data[entity]; ok {
		if err := json.NewEncoder(w).Encode(entitiesMap); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Entity Type not found", http.StatusNotFound)
		return
	}
}

func GetEntity(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {
		if resourcesMap, ok := data[entity]; ok {
			if resourceMap, ok2 := resourcesMap[id]; ok2 {
				if err := json.NewEncoder(w).Encode(resourceMap); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
			http.Error(w, "Entity not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Entity Type not found", http.StatusNotFound)
		return
	}
}

func UpdateEntity(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {

		decoder := json.NewDecoder(r.Body)
		var t map[string]interface{}
		if err := decoder.Decode(&t); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if resourcesMap, ok := data[entity]; ok {
			if resourceMap, ok := resourcesMap[id]; !ok {
				http.Error(w, "Entity Not Found", http.StatusNotFound)
				return
			} else {
				log.Println("Updating Entity")
				for k, v := range t {
					resourceMap[k] = v
				}
				log.Println(resourceMap)

				if err := json.NewEncoder(w).Encode(t); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

	}
}

func DeleteEntity(w http.ResponseWriter, r *http.Request) {
	entity := mux.Vars(r)[PATH_PARAM_ENTITY_TYPE]
	if id, err := strconv.Atoi(mux.Vars(r)[PATH_PARAM_ID]); err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	} else {
		if resourcesMap, ok := data[entity]; ok {
			delete(resourcesMap, id)
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			http.Error(w, "Entity Type not found", http.StatusNotFound)
			return
		}
	}
}
