package models

import "github.com/google/uuid"

type RestRoom struct {
  Id        uuid.UUID   `json:"id"`
  Players   []RestPlayer    `json:"players"`
  Owner     RestPlayer      `json:"owner"`
}
