package models

import "github.com/google/uuid"

type Player struct {
  Id    uuid.UUID `json:"id"`
  Name  string    `json:"name"`
}
