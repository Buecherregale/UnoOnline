package controller

import (
	"encoding/json"
	"net/http"
	"slices"
	"uno_online/game"
	"uno_online/models"
	"uno_online/util"

	"github.com/google/uuid"
)

// POST: /room/
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var pId uuid.UUID
	err := json.NewDecoder(r.Body).Decode(&pId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	owner := game.Players[pId]
	if owner == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	rId := uuid.New()
	room := models.Room{Id: rId, Players: []models.Player{*owner}, Owner: *owner}

	game.Rooms[rId] = &room

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

// POST: /room/{id}/players
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var jId uuid.UUID
	err := json.NewDecoder(r.Body).Decode(&jId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	rId, err := util.ExtractUrlParam(r.URL.Path, 2)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	asUUID, err := uuid.Parse(rId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	room := game.Rooms[asUUID]
	if room == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	joining := game.Players[jId]
	if joining == nil {
		http.Error(w, "Bad reqeust", http.StatusBadRequest)
		return
	}

	room.Players = append(room.Players, *joining)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(room.Id)
}

// Delete: /room/{id}/players/
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var lId uuid.UUID
	err := json.NewDecoder(r.Body).Decode(&lId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	rId, err := util.ExtractUrlParam(r.URL.Path, 2)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	asUUID, err := uuid.Parse(rId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	room := game.Rooms[asUUID]
	if room == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	leaving := game.Players[lId]
	if leaving == nil {
		http.Error(w, "Bad reqeust", http.StatusBadRequest)
		return
	}

	leaveIndex := slices.Index(room.Players, *leaving)
	if leaveIndex == -1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if leaving.Id == room.Owner.Id {
		if len(room.Players) > 1 {
			next := leaveIndex + 1
			next = next % len(room.Players)
			room.Owner = room.Players[next]
		} else {
			game.Rooms[room.Id] = nil
			return
		}
	}

	room.Players = append(room.Players[:leaveIndex], room.Players[leaveIndex+1:]...)
}

// POST: /room/{id}/
func Start(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var pId uuid.UUID
	err := json.NewDecoder(r.Body).Decode(&pId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	rId, err := util.ExtractUrlParam(r.URL.Path, 2)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	asUUID, err := uuid.Parse(rId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	room := game.Rooms[asUUID]
	if room == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if pId != room.Owner.Id {
		http.Error(w, "Not the owner", http.StatusForbidden)
		return
	}

	// start via websocket
}
