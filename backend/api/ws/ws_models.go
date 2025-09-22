package ws

import (
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

const PLAYER_SINGLE_CHANNEL_BUFFER_SIZE int = 7			// all 7 starting card messages 
const PLAYER_BROADCAST_CHANNEL_BUFFER_SIZE int = 3
const PLAYER_RESPOND_CHANNEL_BUFFER_SIZE int = 1

type Message struct {
	Type      		string          `json:"type"`
	Payload   		json.RawMessage `json:"payload"`
	MessageId 		uuid.UUID      	`json:"message_id,omitempty"`
	ExpectsReply 	bool						`json:"expects_reply"`
}

type MsgReceiver func(roomId, playerId uuid.UUID, msg Message)

type WsConn interface {
	ReadJSON(v any) error
	WriteJSON(v any) error
	Close() error
}

type WsPlayer struct {
	id           	uuid.UUID
	conn         	WsConn
	singleChan    chan Message
	broadcastChan	chan Message
	responseChans map[uuid.UUID]chan Message
	mutex        	sync.Mutex
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
