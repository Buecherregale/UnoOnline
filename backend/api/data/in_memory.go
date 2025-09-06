package data

import (
	"uno_online/api/dtos"

	"github.com/google/uuid"
)

var Players = make(map[uuid.UUID]*dtos.Player)
var Rooms = make(map[uuid.UUID]*dtos.Room)
