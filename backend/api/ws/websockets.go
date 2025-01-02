package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Message struct {
	Msg string `json:"msg"`
}

type MsgReceiver func(roomId, playerId uuid.UUID, msg Message)

type WsPlayer struct {
	id       uuid.UUID
	conn     *websocket.Conn
	sendChan chan Message
}

type WsRoom struct {
	id        uuid.UUID
	broadcast chan Message
	players   map[uuid.UUID]*WsPlayer
	handler   MsgReceiver
	mutex     sync.Mutex
}

type Server struct {
	rooms map[uuid.UUID]*WsRoom
	mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		rooms: make(map[uuid.UUID]*WsRoom),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnectMsg(w http.ResponseWriter, r *http.Request, server *Server) {
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
	server.HandleConnection(w, r, roomId, playerId)
}

func (s *Server) BroadcastMsg(roomId uuid.UUID, msg Message) {
	s.mutex.Lock()
	room, exists := s.rooms[roomId]
	s.mutex.Unlock()
	if !exists {
		log.Printf("Room %s does not exist\n", roomId)
		return
	}

	room.broadcast <- msg
}

func (s *Server) SendMsg(roomId, playerId uuid.UUID, msg Message) {
	s.mutex.Lock()
	room, exists := s.rooms[roomId]
	s.mutex.Unlock()
	if !exists {
		log.Printf("Room %s does not exist\n", roomId)
		return
	}

	room.mutex.Lock()
	client, exists := room.players[playerId]
	room.mutex.Unlock()
	if !exists {
		log.Printf("Client %s not found in room %s\n", playerId, roomId)
		return
	}

	client.sendChan <- msg
}

func (s *Server) ReceiveMsg(roomId uuid.UUID, handler MsgReceiver) {
	s.mutex.Lock()
	room, exists := s.rooms[roomId]
	s.mutex.Unlock()
	if !exists {
		log.Printf("Room %s does not exist\n", roomId)
		return
	}

	room.mutex.Lock()
	room.handler = handler
	room.mutex.Unlock()
}

func (s *Server) HandleConnection(w http.ResponseWriter, r *http.Request, roomId, playerId uuid.UUID) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	s.mutex.Lock()
	room, exists := s.rooms[roomId]
	if !exists {
		log.Printf("Room %s does not exist\n", roomId)
		return
	}
	s.mutex.Unlock()

	player := &WsPlayer{
		id:       playerId,
		conn:     conn,
		sendChan: make(chan Message),
	}

	room.AddClient(player)
	go player.ReadMessages(room)
	go player.WriteMessages()
}

func (s *Server) CreateRoom(roomId uuid.UUID, receiver MsgReceiver) {
	s.mutex.Lock()
	if s.rooms[roomId] != nil {
		log.Printf("Room %s already exists\n", roomId)
		return
	}

	room := &WsRoom{
		id:        roomId,
		players:   make(map[uuid.UUID]*WsPlayer),
		broadcast: make(chan Message),
		handler:   receiver,
	}

	s.rooms[roomId] = room
	go room.Run()
	s.mutex.Unlock()
}

func (room *WsRoom) AddClient(player *WsPlayer) {
	room.mutex.Lock()
	room.players[player.id] = player
	room.mutex.Unlock()
}

func (room *WsRoom) RemoveClient(player *WsPlayer) {
	room.mutex.Lock()
	delete(room.players, player.id)
	close(player.sendChan)
	room.mutex.Unlock()
}

func (room *WsRoom) Run() {
	for msg := range room.broadcast {
		room.mutex.Lock()
		for _, client := range room.players {
			client.sendChan <- msg
		}
		room.mutex.Unlock()
	}
}

func (player *WsPlayer) ReadMessages(room *WsRoom) {
	defer func() {
		room.RemoveClient(player)
		player.conn.Close()
	}()

	for {
		var msg Message
		err := player.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("got message to read %s\n", msg)
		if room.handler != nil {
			room.handler(room.id, player.id, msg)
		}
	}
}

func (player *WsPlayer) WriteMessages() {
	defer player.conn.Close()

	for msg := range player.sendChan {
		err := player.conn.WriteJSON(msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
