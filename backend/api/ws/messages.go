package ws

import (
	"encoding/json"

	"uno_online/api/dtos"
)

type ErrorPayload struct {
	Code    	int    				`json:"code,omitempty"`
	Message 	string 				`json:"message,omitempty"`
}

type GameStartPayload struct {
	TopCard 	any					 	`json:"top_card,omitempty"`
}

type CardPlayedPayload struct {
	Player    dtos.Player   `json:"player,omitempty"`
	Card     	any 					`json:"card,omitempty"`
}

type PlayerTurnPayload struct {
	Player    dtos.Player   `json:"player,omitempty"`
}

type PlayerWinPayload struct {
	Player    dtos.Player   `json:"player,omitempty"`
}

type PlayerDrawsCardsPayload struct {
	Player    dtos.Player   `json:"player,omitempty"`
	Amount   	int       		`json:"amount,omitempty"`
}

type PlayerSkippedPayload struct{}

type DirectionChangedPayload struct {
	Direction int 		  		`json:"direction,omitempty"`
}

type RoomJoinPayload struct {
	Player 	 	dtos.Player		`json:"player,omitempty"`
}

type RoomLeftPayload struct {
	Player 		dtos.Player		`json:"player,omitempty"`
	Owner			dtos.Player		`json:"owner,omitempty"`
}

type RoomStartPayload struct { 
	Players  	[]dtos.Player `json:"players,omitempty"`
	Direction int 					`json:"direction,omitempty"`
}

type PlayerChoseColorPayload struct {
	Player    dtos.Player   `json:"player,omitempty"`
	Color    	int       		`json:"color,omitempty"`
}

type AskColorPayload struct {
	Options 	[]int 				`json:"options,omitempty"`
}

type AnswerColorPayload struct {
	Chosen 		int 					`json:"chosen,omitempty"`
}

type AskCardPayload struct {
	Options 	[]any					`json:"options,omitempty"`
}

type AnswerCardPayload struct {
	Card 			any 					`json:"card,omitempty"`
}

type YouDrawCardPayload struct {
	Cards 		[]any 				`json:"cards,omitempty"`
}

func MsgToPayload[T any](message Message) (*T, bool) {
	var payload T
	if err := json.Unmarshal(message.Payload, payload); err != nil {
		return &payload, true
	}
	return nil, false
}
