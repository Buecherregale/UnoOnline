package ws

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/Buecherregale/log"
)

func (room *WsRoom) BroadcastMessage(msgType string, payload any) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Failed to marshal payload: %v\n", err)
		return
	}

	message := Message{
		Type:    msgType,
		Payload: json.RawMessage(payloadBytes),
		ExpectsReply: false,
		MessageId: uuid.New(),
	}
	log.Debugf("Broadcasting message '%s' of type '%s' to room '%s'\n", message.MessageId, message.Type, room.id)

	for _, player := range room.Players {
		player.mutex.Lock()
		select {
		case player.broadcastChan <- message:
		default:
			log.Errorf("Failed to send broadcast message '%s' to player %s: send channel is full or closed\n", message.MessageId, player.id)
		}
		player.mutex.Unlock()
	}
}

func (room *WsRoom) RemovePlayer(playerId uuid.UUID) {
	room.mutex.Lock()
	log.Debugf("Removing player '%s' from room\n", playerId.String())
	player, exists := room.Players[playerId]
	if !exists {
		log.Errorf("Requested player %s does not exist in room %s\n", playerId, room.id)
		return
	}
	log.Debugf("Closing connection to player: %s\n", playerId)
	delete(room.Players, playerId)
	player.Close()
	room.mutex.Unlock()
}

func (room *WsRoom) AddPlayer(player *WsPlayer) {
	log.Debugf("Added player '%s' to WsRoom\n", player.id)
	room.mutex.Lock()
	room.Players[player.id] = player
	room.mutex.Unlock()
}

