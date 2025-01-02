package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"uno_online/game"
	"uno_online/models"
	"uno_online/util"

	"github.com/google/uuid"
)

// struct to unmarshal single uuids
type uuidJson struct {
	Id uuid.UUID `json:"id"`
}

// POST: /room/
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var pId uuidJson
	err := json.NewDecoder(r.Body).Decode(&pId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	owner := game.Players[pId.Id]
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
	var jId uuidJson
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

	joining := game.Players[jId.Id]
	if joining == nil {
		http.Error(w, "Bad reqeust", http.StatusBadRequest)
		return
	}
	if slices.Contains(room.Players, *joining) {
		http.Error(w, "Already joined", http.StatusBadRequest)
		return
	}

	room.Players = append(room.Players, *joining)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(room)
}

// Delete: /room/{id}/players/
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	var lId uuidJson
	err := json.NewDecoder(r.Body).Decode(&lId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	rId, err := util.ExtractUrlParam(r.URL.Path, 2)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	asUUID, err := uuid.Parse(rId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	room := game.Rooms[asUUID]
	if room == nil {
		fmt.Println(err)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	leaving := game.Players[lId.Id]
	if leaving == nil {
		fmt.Println(err)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	leaveIndex := slices.Index(room.Players, *leaving)
	if leaveIndex == -1 {
		http.Error(w, "Already left", http.StatusBadRequest)
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
	var pId uuidJson
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

	if pId.Id != room.Owner.Id {
		http.Error(w, "Not the owner", http.StatusForbidden)
		return
	}

	// start via websocket
}

// GET: /room/{id}/
func GetRoom(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}
