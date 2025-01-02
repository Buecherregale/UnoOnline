package main

import (
	"net/http"
	"uno_online/controller"
)

func main() {
	mux := Router()
	http.ListenAndServe(":8080", mux)
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /player", controller.CreatePlayer)
	mux.HandleFunc("POST /room", controller.CreateRoom)
	mux.HandleFunc("POST /room/{rId}", controller.Start)
	mux.HandleFunc("POST /room/{rId}/players", controller.JoinRoom)
	mux.HandleFunc("DELETE /room/{rId}/players", controller.LeaveRoom)

	return mux
}
