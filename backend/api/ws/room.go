package ws

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

func (room *WsRoom) BroadcastMessage(msgType string, payload interface{}) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal payload: %v\n", err)
		return
	}

	message := Message{
		Type:    msgType,
		Payload: json.RawMessage(payloadBytes),
	}

	room.mutex.Lock()
	defer room.mutex.Unlock()

	for _, player := range room.players {
		select {
		case player.sendChan <- message:
		default:
			log.Printf("failed to send message to player %s: send channel is full or closed\n", player.id)
		}
	}
}

func (room *WsRoom) RemovePlayer(playerId uuid.UUID) {
	room.mutex.Lock()
	player, exists := room.players[playerId]
	if !exists {
		log.Printf("requested player %s does not exist in room %s\n", playerId, room.id)
		return
	}
	player.conn.Close()
	delete(room.players, playerId)
	close(player.sendChan)
	room.mutex.Unlock()
}

func (room *WsRoom) AddPlayer(player *WsPlayer) {
	room.mutex.Lock()
	room.players[player.id] = player
	room.mutex.Unlock()
}

func (room *WsRoom) Run() {
	for msg := range room.broadcast {
		room.mutex.Lock()
		for _, player := range room.players {
			player.sendChan <- msg
		}
		room.mutex.Unlock()
	}
}
