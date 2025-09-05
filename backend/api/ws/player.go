package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func (player *WsPlayer) SendMessage(msgType string, payload any) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("could not marshal message %s\n", payload)
		return
	}
	msg := Message{
		Type:    msgType,
		Payload: json.RawMessage(bytes),
	}

	select {
	case player.sendChan <- msg:
		return
	default:
		log.Printf("failed to send message to player %s: send channel is full or closed\n", player.id)
	}
}

func (player *WsPlayer) AskAndWaitReply(msgType string, payload any, timeout time.Duration) (*Message, bool, error) {
	messageId := uuid.New()

	responseChan := make(chan Message, 1)
	player.mutex.Lock()
	if player.responseChan == nil {
		player.responseChan = make(map[uuid.UUID]chan Message)
	}
	player.responseChan[messageId] = responseChan
	player.mutex.Unlock()

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	message := Message{
		Type:      msgType,
		MessageId: &messageId,
		Payload:   json.RawMessage(payloadBytes),
	}

	player.sendChan <- message

	select {
	case response := <-responseChan:
		return &response, false, nil
	case <-time.After(timeout):
		// clean up response channel
		player.mutex.Lock()
		delete(player.responseChan, messageId)
		player.mutex.Unlock()
		return nil, true, fmt.Errorf("timeout waiting for response")
	}
}

func (player *WsPlayer) SendError(code int, msg string) {
	err := ErrorPayload{
		Code:    code,
		Message: msg,
	}
	player.SendMessage("ErrorMessage", err)
}
