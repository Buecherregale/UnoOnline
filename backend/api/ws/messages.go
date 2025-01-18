package ws

import "github.com/google/uuid"

type CardPayload struct {
	Value string `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

type ErrorPayload struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type CardPlayedPayload struct {
	PlayerId uuid.UUID   `json:"player_id,omitempty"`
	Card     CardPayload `json:"card,omitempty"`
}

type PlayerTurnPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
}

type PlayerWinPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
}

type PlayerDrawsCardsPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Amount   int       `json:"amount,omitempty"`
}

type PlayerSkippedPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
}

type PlayerChangesDirectionPayload struct {
	PlayerId  uuid.UUID `json:"player_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Direction int       `json:"direction,omitempty"`
}

type PlayerChosesColorPayload struct {
	PlayerId uuid.UUID `json:"player_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Color    string    `json:"color,omitempty"`
}

type AskColorPayload struct {
	Options []string `json:"options,omitempty"`
}

type AnswerColorPayload struct {
	Chosen string `json:"chosen,omitempty"`
}

type AskCardPayload struct {
	Options []CardPayload `json:"options,omitempty"`
}

type AnswerCardPayload struct {
	Card CardPayload `json:"card,omitempty"`
}

type YouDrawCardPayload struct {
	Cards []CardPayload `json:"cards,omitempty"`
}
