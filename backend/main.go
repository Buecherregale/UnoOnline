package main

import (
	"net/http"
	"uno_online/controller"
	"uno_online/game"
	"uno_online/models"

	"github.com/google/uuid"
)

func main() {
	FillTestData()

	mux := Router()
	http.ListenAndServe(":8080", mux)
}

func FillTestData() {
	p1iD := uuid.MustParse("e5384075-99f9-474e-85d7-0bb4bb0c62a7")
	p2iD := uuid.MustParse("52041871-8abf-4d03-8615-349753a791b6")
	p3iD := uuid.MustParse("dd4d2e03-d670-4097-bdec-22cecfdbdefc")
	p1 := models.Player{Name: "Klaus", Id: p1iD}
	p2 := models.Player{Name: "Biggie Smalls", Id: p2iD}
	p3 := models.Player{Name: "Wilhelm", Id: p3iD}
	game.Players[p1iD] = &p1
	game.Players[p2iD] = &p2
	game.Players[p3iD] = &p3

	room1iD := uuid.MustParse("4d3e97bf-cc2e-4af0-9397-2a0e3b331c6f")
	game.Rooms[room1iD] = &models.Room{Id: room1iD, Owner: p1, Players: []models.Player{p1}}
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /player", controller.CreatePlayer)
	mux.HandleFunc("POST /room", controller.CreateRoom)
	mux.HandleFunc("POST /room/{rId}", controller.Start)
	mux.HandleFunc("GET /room/{rId}", controller.GetRoom)
	mux.HandleFunc("POST /room/{rId}/players", controller.JoinRoom)
	mux.HandleFunc("DELETE /room/{rId}/players", controller.LeaveRoom)

	return mux
}
