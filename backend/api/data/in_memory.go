package data

import (
	"uno_online/api/models"

	"github.com/google/uuid"
)

var Players = make(map[uuid.UUID]*models.Player)
var Rooms = make(map[uuid.UUID]*models.Room)
