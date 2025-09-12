package ws_test

import (
	"testing"
	"time"
	"uno_online/api/dtos"
	"uno_online/api/ws"

	"github.com/google/uuid"
)

func TestBroadcastMessage(t *testing.T) {
	tested := ws.MockWsRoom()
	c1 := ws.MockWsConn()
	ws.MockWsPlayer(c1, tested)
	c2 := ws.MockWsConn()
	ws.MockWsPlayer(c2, tested)

	expectedType := "PlayerChoseColorPayload"
	tested.BroadcastMessage(expectedType, ws.PlayerChoseColorPayload {
		Player: dtos.Player{ 
			Id: uuid.New(), 
			Name: "dto player", 
		},
		Color: 0,
	})

	time.Sleep(10 * time.Millisecond)

	if len(c1.Written) != 1 {
		t.Fatalf("p1 expected 1 message, got %d", len(c1.Written))
	}
	if len(c2.Written) != 1 {
		t.Fatalf("p2 expected 1 message, got %d", len(c2.Written))
	}
	if c1.Written[0].Type != expectedType || c2.Written[0].Type != expectedType {
		t.Fatalf("expected type %q, got %q", expectedType, c1.Written[0].Type)
	}
}

func TestAddPlayer(t *testing.T) {
	c := ws.MockWsConn()
	p := ws.MockWsPlayer(c, nil)
	tested := ws.MockWsRoom()
	tested.AddPlayer(p)

	if len(tested.Players) != 1 {
		t.Fatalf("expected 1 player in room, got %d", len(tested.Players))
	}
}

