package models

import "github.com/google/uuid"

type Room struct {
  Id        uuid.UUID   `json:"id"`
  Players   []uuid.UUID `json:"players"`
  Owner     []uuid.UUID `json:"owner"`
}
