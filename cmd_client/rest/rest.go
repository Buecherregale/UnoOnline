package rest

import (
	"bufio"
	"bytes"
	"cmd_client/models"
	"cmd_client/websocket"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func SendRestRequest[T any](method, endpoint string, query map[string]string, body any, replyBody T, hasReply bool) *T {
	baseUrl := "http://localhost:8080" + endpoint
	params := url.Values{}
	for k, v := range query {
		params.Add(k, v)
	}
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	encBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("couldn't send (marshal) %v\n", err)
		return nil
	}
	req, err := http.NewRequest(method, fullUrl, bytes.NewBuffer(encBody))
	if err != nil {
		fmt.Printf("couldn't send (req) %v\n", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("cound't send (do) %v\n", err)
		return nil
	}
	defer resp.Body.Close()
	fmt.Printf("response: %d\n", resp.StatusCode)
	if !hasReply {
		return nil
	}
	err = json.NewDecoder(resp.Body).Decode(&replyBody)
	if err != nil {
		fmt.Printf("couldn't send (decode) %v\n", err)
		fmt.Printf("status: %s\n", resp.Status)
		return nil
	}
	return &replyBody
}

type uuidJson struct {
	Id uuid.UUID `json:"id,omitempty"`
}
type nameJson struct {
	Name string `json:"name,omitempty"`
}

func JoinLobby(scanner *bufio.Reader) (*models.RestPlayer, *models.RestRoom) {
	var player *models.RestPlayer
	var room *models.RestRoom

	fmt.Println("Enter player name: ")
	name, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(name)
	var resp models.RestPlayer
	player = SendRestRequest("POST", "/player", make(map[string]string), nameJson{Name: name}, resp, true)

	fmt.Printf("created player with name %s and id %s\n", player.Name, player.Id)

	fmt.Println("type 'join' or 'create' to join or create a room")
	cmd, _ := scanner.ReadString('\n')
	cmd = strings.ToLower(cmd)
	if cmd == "join" {
		fmt.Println("enter room uuid:")
		rIds, _ := scanner.ReadString('\n')
		rId := uuid.MustParse(rIds)
		var resp models.RestRoom
		room = SendRestRequest("POST", fmt.Sprintf("/room/%s/players", rId), make(map[string]string), uuidJson{Id: player.Id}, resp, true)
	} else {
		fmt.Println("creating room")
		var resp models.RestRoom
		room = SendRestRequest("POST", "/room", make(map[string]string), uuidJson{Id: player.Id}, resp, true)
	}
	fmt.Printf("now part of room: %s\n", room)
	fmt.Println("now connecting to websocket")
	ws := websocket.JoinRoom(room, player)
	fmt.Println(ws)
	fmt.Println("type start to start (as owner)")
	cmd, _ = scanner.ReadString('\n')
	cmd = strings.ToLower(strings.TrimSpace(cmd))
	if cmd == "start" {
		SendRestRequest("POST", fmt.Sprintf("/room/%s", room.Id), make(map[string]string), uuidJson{Id: player.Id}, models.RestPlayer{}, false)
		fmt.Println("started room")
	}
	return player, room
}
