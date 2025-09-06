package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"sync"

	"uno_online/api/data"
	"uno_online/api/dtos"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type      string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
	MessageId *uuid.UUID      `json:"message_id,omitempty"`
}

type MsgReceiver func(roomId, playerId uuid.UUID, msg Message)

type WsPlayer struct {
	id           uuid.UUID
	conn         *websocket.Conn
	sendChan     chan Message
	responseChan map[uuid.UUID]chan Message
	mutex        sync.Mutex
}

type WsRoom struct {
	id        uuid.UUID
	broadcast chan Message
	Players   map[uuid.UUID]*WsPlayer
	handler   MsgReceiver
	mutex     sync.Mutex
}

type WsServer struct {
	Rooms map[uuid.UUID]*WsRoom
	mutex sync.Mutex
}

func NewServer() *WsServer {
	return &WsServer{
		Rooms: make(map[uuid.UUID]*WsRoom),
	}
}

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
		log.Printf("Room %s already exists\n", roomId)
		return nil
	}

	if data.Rooms[roomId] == nil {
		log.Printf("Room %s does not exist yet. Firstly create via rest\n", roomId)
		return nil
	}

	room := &WsRoom{
		id:        roomId,
		Players:   make(map[uuid.UUID]*WsPlayer),
		broadcast: make(chan Message),
		handler:   receiver,
	}

	s.Rooms[roomId] = room
	go room.Run()
	s.mutex.Unlock()

	return room
}

func (s *WsServer) handleConnection(w http.ResponseWriter, r *http.Request, roomId, playerId uuid.UUID) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	s.mutex.Lock()
	room, exists := s.Rooms[roomId]
	if !exists {
		log.Printf("Room %s does not exist\n", roomId)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	s.mutex.Unlock()

	ro := data.Rooms[roomId]
	if !slices.ContainsFunc(ro.Players, func(p dtos.Player) bool { return p.Id == playerId }) {
		log.Printf("Player not in room\n")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	player := &WsPlayer{
		id:       playerId,
		conn:     conn,
		sendChan: make(chan Message),
	}

	room.AddPlayer(player)
	go player.readMessages(room)
	go player.writeMessages()
}

func (player *WsPlayer) readMessages(room *WsRoom) {
	defer func() {
		room.RemovePlayer(player.id)
		player.conn.Close()
	}()

	for {
		var msg Message
		err := player.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Check if the message has a RequestID and handle it
		if msg.MessageId != nil {
			player.mutex.Lock()
			responseChan, exists := player.responseChan[*msg.MessageId]
			if exists {
				responseChan <- msg
				close(responseChan)
				delete(player.responseChan, *msg.MessageId)
				player.mutex.Unlock()
				continue
			}
			player.mutex.Unlock()
		}

		// Otherwise, process the message normally
		log.Printf("Received message: %+v\n", msg)
		if room.handler != nil {
			room.handler(room.id, player.id, msg)
		}
	}
}

func (player *WsPlayer) writeMessages() {
	defer player.conn.Close()

	for msg := range player.sendChan {
		err := player.conn.WriteJSON(msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
