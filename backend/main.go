package main

import (
	"net/http"
	"time"

	"uno_online/api/controller"
	"uno_online/api/data"
	"uno_online/api/dtos"
	"uno_online/api/ws"

	"github.com/Buecherregale/log"
	"github.com/google/uuid"
)

func main() {
	config := log.LogConfig {
		Level: log.LEVEL_DEBUG,
		Timeformat: time.RFC3339,
		SerializationStrategy: log.STRATEGY_SIMPLE,
		TargetMode: log.TARGET_STDOUT,
		Logfile: "",
	}
	log.Init(config)

	FillTestData()

	mux := Router()

	FillWsTestData(ws.Server)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.HandleConnectMsg(w, r, ws.Server)
	})

	log.Infoln("Starting server on port 8080...")
	log.Fatalf("%v\n", http.ListenAndServe(":8080", mux))
}

func FillTestData() {
	// data to allow for easy testing with curl
	p1iD := uuid.MustParse("e5384075-99f9-474e-85d7-0bb4bb0c62a7")
	p2iD := uuid.MustParse("52041871-8abf-4d03-8615-349753a791b6")
	p3iD := uuid.MustParse("dd4d2e03-d670-4097-bdec-22cecfdbdefc")
	p1 := dtos.Player{Name: "Klaus", Id: p1iD}
	p2 := dtos.Player{Name: "Biggie Smalls", Id: p2iD}
	p3 := dtos.Player{Name: "Wilhelm", Id: p3iD}
	data.Players[p1iD] = &p1
	data.Players[p2iD] = &p2
	data.Players[p3iD] = &p3

	room1iD := uuid.MustParse("4d3e97bf-cc2e-4af0-9397-2a0e3b331c6f")
	data.Rooms[room1iD] = &dtos.Room{Id: room1iD, Owner: p1, Players: []dtos.Player{p1}}
}

func FillWsTestData(server *ws.WsServer) {
	// data to allow for easy testing e.g. using websocat
	room1iD := uuid.MustParse("4d3e97bf-cc2e-4af0-9397-2a0e3b331c6f")
	wsRoom, _ := server.Rooms[room1iD]

	receiver := func(roomId, playerId uuid.UUID, msg ws.Message) {
		log.Debugf("Message send to room %s\nby player %s:\n%s", roomId, playerId, msg.Type)
		wsRoom.BroadcastMessage(msg.Type, msg.Payload)
		wsRoom.Players[playerId].SendMessage(msg.Type, msg.Payload)
	}
	server.CreateRoom(room1iD, receiver)
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /players", controller.CreatePlayer)
	mux.HandleFunc("POST /rooms", controller.CreateRoom)
	mux.HandleFunc("POST /rooms/{rId}", controller.StartRoom)
	mux.HandleFunc("POST /rooms/{rId}/startgame", controller.StartGame)
	mux.HandleFunc("GET /rooms/{rId}", controller.GetRoom)
	mux.HandleFunc("POST /rooms/{rId}/players", controller.JoinRoom)
	mux.HandleFunc("DELETE /rooms/{rId}/players", controller.LeaveRoom)

	return mux
}
