package ws

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ErrorPayload struct {
	Code    int    			`json:"code,omitempty"`
	Message string 			`json:"message,omitempty"`
}

type GameStartPayload struct {
	TopCard any		 			`json:"top_card,omitempty"`
}

type CardPlayedPayload struct {
	PlayerId uuid.UUID   `json:"player_id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Card     any				 `json:"card,omitempty"`
}

type PlayerTurnPayload struct {
	PlayerId uuid.UUID 	`json:"player_id,omitempty"`
	Name     string  		`json:"name,omitempty"`
}

type PlayerWinPayload struct {
	PlayerId uuid.UUID 	`json:"player_id,omitempty"`
	Name     string    	`json:"name,omitempty"`
}

type PlayerDrawsCardsPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Amount   int       `json:"amount,omitempty"`
}

type PlayerSkippedPayload struct{}

type DirectionChangedPayload struct {
	Direction int 		`json:"direction,omitempty"`
}

type PlayerChoseColorPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Color    int       `json:"color,omitempty"`
}

type AskColorPayload struct {
	Options []int 		`json:"options,omitempty"`
}

type AnswerColorPayload struct {
	Chosen int 				`json:"chosen,omitempty"`
}

type AskCardPayload struct {
	Options []any		 `json:"options,omitempty"`
}

type AnswerCardPayload struct {
	Card any				 `json:"card,omitempty"`
}

type YouDrawCardPayload struct {
	Cards []any			 `json:"cards,omitempty"`
}

func MsgToPayload[T any](message Message) (*T, bool) {
	var payload T
	if err := json.Unmarshal(message.Payload, payload); err != nil {
		return &payload, true
	}
	return nil, false
}
