package websocket

import (
	"cmd_client/models"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func JoinRoom(room *models.RestRoom, player *models.RestPlayer) *websocket.Conn {
	queryParams := url.Values{}
	queryParams.Add("roomId", room.Id.String())
	queryParams.Add("playerId", player.Id.String())

	// Add query parameters to the URL
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws", RawQuery: queryParams.Encode()}

	fmt.Printf("connecting to ws %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("error connecting to websocket:", err)
		return nil
	}

	fmt.Println("connected to websocket")
	defer conn.Close()
	fmt.Println("Connected to WebSocket server")

	// Handle interrupts to clean up the connection properly
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Set up a ticker to send ping messages to keep the connection alive
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	done := make(chan struct{})

	// Read messages from the server
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Main loop to keep the connection alive
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Connection closed by server")
				return
			case <-ticker.C:
				// Send a ping message to keep the connection alive
				err := conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					log.Printf("Error sending ping: %v", err)
					return
				}
			case <-interrupt:
				fmt.Println("Interrupt received, closing connection")
				// Send a close message to the server
				err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Printf("Error sending close message: %v", err)
				}
				return
			}
		}
	}()
	return conn
}
