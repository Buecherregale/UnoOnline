package models

import "github.com/google/uuid"

type RestPlayer struct {
  Id uuid.UUID `json:"id"`
  Name string `json:"name"`
}
