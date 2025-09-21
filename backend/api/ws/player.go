package ws

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Buecherregale/log"
	"github.com/google/uuid"
)

// sends a message to the player singleChan, NOT expecting a reply
func (player *WsPlayer) SendMessage(msgType string, payload any) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Could not marshal message %s\n", payload)
		return
	}
	msg := Message{
		Type:    msgType,
		Payload: json.RawMessage(bytes),
		MessageId: uuid.New(),
		ExpectsReply: false,
	}
	log.Debugf("Sending message '%s' of type '%s' to player '%s'\n", msg.MessageId, msgType, player.id)
	player.mutex.Lock()
	defer player.mutex.Unlock()
	select {
	case player.singleChan <- msg:
	default:
		log.Errorf("Failed to send message '%s' to player %s: send channel is full or closed\n", msg.MessageId, player.id)
	}
}

func (player *WsPlayer) AskAndWaitReply(msgType string, payload any, timeout time.Duration) (*Message, bool, error) {
	messageId := uuid.New()
	log.Debugf("Asking player '%s' for reply to message '%s' of type: '%s'\n", player.id, messageId, msgType)

	responseChan := make(chan Message, 1)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, false, err
	}

	message := Message{
		Type:      msgType,
		MessageId: messageId,
		ExpectsReply: true,
		Payload:   json.RawMessage(payloadBytes),
	}
	
	player.mutex.Lock()
	player.responseChans[messageId] = responseChan
	player.singleChan <- message
	player.mutex.Unlock()

	select {
	case response := <-responseChan:
		// cleanup here probably necessary too
		return &response, false, nil
	case <-time.After(timeout):
		// clean up response channel
		player.mutex.Lock()
		delete(player.responseChans, messageId)
		player.mutex.Unlock()
		return nil, true, fmt.Errorf("timeout waiting for response")
	}
}

// perform cleanup on the player, closing all channels and connections
func (player *WsPlayer) Close() {
	log.Debugf("Closing player: %s\n", player.id)
	player.mutex.Lock()
	defer player.mutex.Unlock()
	for id, c := range player.responseChans {
		close(c)
		delete(player.responseChans, id)
	}
	close(player.singleChan)
	close(player.broadcastChan)
	player.conn.Close()
}

func (player *WsPlayer) SendError(code int, msg string) {
	err := ErrorPayload{
		Code:    code,
		Message: msg,
	}
	log.Debugf("Sending error to player '%s': %+v\n", player.id, err)
	player.SendMessage("ErrorMessage", err)
}

// read player messages from player.conn, writing them to the player.responseChan if they have an id
func (player *WsPlayer) readMessages(room *WsRoom) {
	defer func() {
		room.RemovePlayer(player.id)
		log.Debugf("Done reading messages. Closing connection to player: %s\n", player.id)
		player.conn.Close()
	}()

	for {
		var msg Message
		err := player.conn.ReadJSON(&msg)
		if err != nil {
			log.Errorf("Read error: %v\n", err)
			break
		}

		// Check if the message has a RequestID and handle it
		if msg.ExpectsReply {
			player.mutex.Lock()
			responseChan, exists := player.responseChans[msg.MessageId]
			if exists {
				responseChan <- msg
				close(responseChan)
				delete(player.responseChans, msg.MessageId)
				player.mutex.Unlock()
				continue
			}
			player.mutex.Unlock()
		}

		// Otherwise, process the message normally
		log.Debugf("Received message: %+v\n", msg)
		if room.handler != nil {
			room.handler(room.id, player.id, msg)
		}
	}
}

// goroutine func to listen to messages on player.singleChan, sending them automatically
func (player *WsPlayer) autoSendSingleMessages() {
	log.Debugf("Starting SINGLE message listener for player: %s\n", player.id)
	defer func() { 
		log.Debugf("Closing connection to player: %s\n", player.id)
		player.conn.Close()
	}()
	for msg := range player.singleChan {
		log.Debugf("Auto send SINGLE message '%s' to player '%s'. Length of channel is: %d\n", msg.MessageId, player.id, len(player.singleChan))
		player._sendMessage(msg)
	}
}

// goroutine func to lsten to messages on player.broadcastChan, sending them automatically
func (player *WsPlayer) autoSendBroadcastMessages() {
	log.Debugf("Starting BROADCAST message listener for player: %s\n", player.id)
	defer func() {
		log.Debugf("Closing connection to player: %s\n", player.id)
		player.conn.Close()
	}()
	for msg := range player.broadcastChan {
		log.Debugf("Auto send BROADCAST message '%s' to player '%s'. Length of channel is: %d\n", msg.MessageId, player.id, len(player.broadcastChan))
		player._sendMessage(msg)
	}
}

// writes the given message to the player.conn. TO SEND MESSAGES TO THE PLAYER USE 'player.SendMessage()'
func (player *WsPlayer) _sendMessage(msg Message) error {
	log.Debugf("Sending message '%s' to player: %s\n", msg.MessageId, player.id)
	player.mutex.Lock()
	defer player.mutex.Unlock()
	err := player.conn.WriteJSON(msg)
	if err != nil {
		log.Errorf("Error sending message '%s' to player '%s': %v\n", msg.MessageId, player.id, err)
	}
	return err
}

