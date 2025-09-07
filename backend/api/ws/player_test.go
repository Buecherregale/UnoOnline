package ws_test

import (
	"testing"
	"time"

	"uno_online/api/dtos"
	"uno_online/api/ws"

	"github.com/google/uuid"
)

func TestPlayerSendMessage(t *testing.T) {
	mockConn := ws.MockWsConn()
	mocked := ws.MockWsPlayer(mockConn, nil)

	expectedType := "RoomJoinPayload"
	expectedPayload := ws.RoomJoinPayload {
		Player: dtos.Player{Id: uuid.New(), Name: "name"},
	}

	mocked.SendMessage(expectedType, expectedPayload)

	time.Sleep(10 * time.Millisecond) // let go routine run

	if len(mockConn.Written) != 1 {
		t.Fatalf("expected 1 message, got %d", len(mockConn.Written))
	}
	if mockConn.Written[0].Type != expectedType {
		t.Fatalf("expected type %q, got %q", expectedType, mockConn.Written[0].Type)
	}
}

func TestPlayerSendError(t *testing.T) {
	conn := ws.MockWsConn()
	player := ws.MockWsPlayer(conn, nil)

	player.SendError(500, "Testing server error")

	time.Sleep(10 * time.Millisecond)

	if len(conn.Written) != 1 {
		t.Fatalf("expected 1 message, got %d", len(conn.Written))
	}
	if conn.Written[0].Type != "ErrorMessage" {
		t.Fatalf("expected type ErrorMessage, got %q", conn.Written[0].Type)
	}
}

func TestPlayerAskAndAwaitReply(t *testing.T) {
	conn := ws.MockWsConn()
	room := ws.MockWsRoom()
	resp := make(map[uuid.UUID]chan ws.Message)
	player := ws.MockWsPlayerResponseChan(conn, room, &resp)

	payload := ws.AskCardPayload {
		Options: []any{"a", "b"},
	}
	player.AskAndWaitReply("AskCardPayload", payload, 1 * time.Second)

	time.Sleep(10 * time.Millisecond)

	if len(conn.Written) != 1 {
		t.Fatalf("expected 1 message, got %d", len(conn.Written))
	}
	if conn.Written[0].Type != "AskCardPayload" {
		t.Fatalf("expected type AskCardPayload, got %q", conn.Written[0].Type)
	}
}

