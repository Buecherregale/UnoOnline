package dtos

import "github.com/google/uuid"

type Room struct {
  Id        uuid.UUID   `json:"id"`
  Players   []Player    `json:"players"`
  Owner     Player      `json:"owner"`
}
