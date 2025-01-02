package controller

import (
	"encoding/json"
	"net/http"

	"uno_online/game"
	"uno_online/models"

	"github.com/google/uuid"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var name string
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	game.Players[id] = &models.Player{Id: id, Name: name}
}
