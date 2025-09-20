package ws

import (
	"io"

	"github.com/Buecherregale/log"
	"github.com/google/uuid"
)

type MockedConn struct {
	Written []Message
	Read chan Message
	Closed bool
}

func MockWsConn() *MockedConn {
	return &MockedConn{
		Read: make(chan Message, 10),
	}
}

func (conn *MockedConn) WriteJSON(v any) error {
	msg := v.(Message)
	conn.Written = append(conn.Written, msg)
	return nil
}

func (conn *MockedConn) ReadJSON(v any) error {
	msg, ok := <-conn.Read 
	if !ok {
		return io.EOF
	}
	*v.(*Message) = msg 
	return nil
}

func (conn *MockedConn) Close() error {
	log.Debugf("closing mocked connection...\n")
	conn.Closed = true 
	close(conn.Read)
	return nil 
}

func MockWsPlayer(conn WsConn, room *WsRoom) *WsPlayer {
	player := WsPlayer{
		id: uuid.New(),
		conn: conn,
		singleChan: make(chan Message, 1),
		broadcastChan: make(chan Message),
	}
	go player.autoSendSingleMessages()
	go player.autoSendBroadcastMessages()
	if room != nil {
		room.AddPlayer(&player)
		go player.readMessages(room)
	}
	return &player
}

func MockWsPlayerResponseChan(conn WsConn, room *WsRoom, responseChan *map[uuid.UUID]chan Message) *WsPlayer {
	player := WsPlayer {
		id: uuid.New(),
		conn: conn, 
		singleChan: make(chan Message, 1),
		broadcastChan: make(chan Message),
		responseChans: *responseChan,
	}
	go player.autoSendSingleMessages()
	go player.autoSendBroadcastMessages()
	if room != nil {
		room.Players[player.id] = &player
		go player.readMessages(room)
	}
	return &player
}

func MockWsRoom() *WsRoom {
	return &WsRoom{
		id: uuid.New(),
		Players: make(map[uuid.UUID]*WsPlayer),
		handler: func(roomId, playerId uuid.UUID, msg Message) {},
	}
}

