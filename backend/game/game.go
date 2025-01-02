package game

import (
	"uno_online/models"

	"github.com/google/uuid"
)

var Players = make(map[uuid.UUID]*models.Player)
var Rooms = make(map[uuid.UUID]*models.Room)
