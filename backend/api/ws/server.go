package ws

import (
	"net/http"
	"slices"

	"uno_online/api/data"
	"uno_online/api/dtos"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/Buecherregale/log"
)

var Server = NewServer()

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnectMsg(w http.ResponseWriter, r *http.Request, server *WsServer) {
	roomId, err := uuid.Parse(r.URL.Query().Get("roomId"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	playerId, err := uuid.Parse(r.URL.Query().Get("playerId"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	server.handleConnection(w, r, roomId, playerId)
}

func (s *WsServer) CreateRoom(roomId uuid.UUID, receiver MsgReceiver) *WsRoom {
	s.mutex.Lock()
	if s.Rooms[roomId] != nil {
		log.Errorf("Room %s already exists\n", roomId)
		return nil
	}

	if data.Rooms[roomId] == nil {
		log.Errorf("Room %s does not exist yet. Firstly create via rest\n", roomId)
		return nil
	}

	room := &WsRoom{
		id:        roomId,
		Players:   make(map[uuid.UUID]*WsPlayer),
		handler:   receiver,
	}
	log.Debugf("Created new WsRoom: %s\n", roomId)

	s.Rooms[roomId] = room
	s.mutex.Unlock()

	return room
}

func (s *WsServer) handleConnection(w http.ResponseWriter, r *http.Request, roomId, playerId uuid.UUID) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Upgrade error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	s.mutex.Lock()
	room, exists := s.Rooms[roomId]
	if !exists {
		log.Errorf("Room %s does not exist\n", roomId)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	s.mutex.Unlock()

	ro := data.Rooms[roomId]
	if !slices.ContainsFunc(ro.Players, func(p dtos.Player) bool { return p.Id == playerId }) {
		log.Errorf("Player not in room\n")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	player := &WsPlayer{
		id:       			playerId,
		conn:     			conn,
		singleChan: 		make(chan Message, PLAYER_SINGLE_CHANNEL_BUFFER_SIZE),
		broadcastChan: 	make(chan Message, PLAYER_BROADCAST_CHANNEL_BUFFER_SIZE),
		responseChans: 	make(map[uuid.UUID]chan Message),
	}
	room.AddPlayer(player)
	
	go player.readMessages(room)
	go player.autoSendSingleMessages()
	go player.autoSendBroadcastMessages()
}

