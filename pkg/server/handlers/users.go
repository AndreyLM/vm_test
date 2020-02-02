package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/andreylm/vm_test/pkg/storage"
)

// UserList - get all users
func UserList(db storage.IStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetAll()
		if err != nil {
			log.Println(err)
			sendResponse(w, map[string]interface{}{
				"status": http.StatusInternalServerError,
				"error":  "Error getting data",
			})
			return
		}

		sendResponse(w, map[string]interface{}{
			"status": http.StatusOK,
			"users":  users,
		})
	}
}

// UserByID - gets user by id
func UserByID(db storage.IStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			sendResponse(w, map[string]interface{}{
				"status": http.StatusBadRequest,
				"error":  "Error getting data",
			})
			return
		}

		u, err := db.Get(id)
		if err != nil {
			log.Println(err)
		}

		sendResponse(w, map[string]interface{}{
			"status": http.StatusOK,
			"user":   u,
		})
	}
}

func sendResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
