package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
)

type Server struct {
	connections map[*websocket.Conn]bool
	inputStream chan []byte
}

func NewServer() *Server {
	return &Server{connections: make(map[*websocket.Conn]bool)}
}

func (s *Server) WsHandler(ws *websocket.Conn) {
	s.connections[ws] = true
	s.read(ws)
}

func (s *Server) read(ws *websocket.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := ws.Read(buffer)
		if err != nil {
			if err == io.EOF {
				delete(s.connections, ws)
				break
			}
			fmt.Println(err)
			continue
		}
		s.inputStream <- buffer[:n]

	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.connections {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println(err)
			}
		}(ws)
	}
}

//http.handle("/",websocket.Handler(server.WsHandler))
