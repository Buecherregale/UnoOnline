package controller

import (
	"encoding/json"
	"net/http"

	"uno_online/api/data"
	"uno_online/api/dtos"

	"github.com/google/uuid"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	player := dtos.Player{Id: id, Name: p.Name}
	data.Players[id] = &player

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}
