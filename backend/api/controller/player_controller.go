package controller

import (
	"encoding/json"
	"net/http"

	"uno_online/api/data"
	"uno_online/api/models"

	"github.com/google/uuid"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var p struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	data.Players[id] = &models.Player{Id: id, Name: p.Name}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uuidJson{Id: id})
}
