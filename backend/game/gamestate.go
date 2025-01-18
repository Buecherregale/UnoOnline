package game

import (
	"uno_online/api/models"
	"uno_online/api/ws"
)

type CardPlayEventListener func(gp *GamePlayer, card *Card, state *GameState)

type GamePlayer struct {
	P    *models.Player
	WsP  *ws.WsPlayer
	Hand []Card
}

type GameState struct {
	Room    *models.Room
	WsRoom  *ws.WsRoom
	Players []*GamePlayer
	Deck    *Deck
	Stack   *Stack
	Winner  *GamePlayer
	CurrDir int
	CurrI   int

	listeners []CardPlayEventListener
}

func (state *GameState) NextPlayer() *GamePlayer {
	state.CurrI += state.CurrDir
	state.CurrI %= len(state.Players)
	next := state.Players[state.CurrI]

	state.WsRoom.BroadcastMessage("PlayerTurnPayload", ws.PlayerTurnPayload{
		PlayerId: next.P.Id,
		Name:     next.P.Name,
	})

	return next
}

func (state *GameState) PeekNextPlayer() *GamePlayer {
	next := state.CurrI + state.CurrDir
	next %= len(state.Players)
	return state.Players[next]
}

func (state *GameState) RegisterListener(listener CardPlayEventListener) {
	state.listeners = append(state.listeners, listener)
}

func (state *GameState) DrawCards(target *GamePlayer, amount int) {
	for range amount {
		card := state.Deck.Draw()
		target.Hand = append(target.Hand, card)
		target.WsP.SendMessage("YouDrawCardPayload", ws.YouDrawCardPayload{
			Cards: []interface{}{card},
		})
	}
	state.WsRoom.BroadcastMessage("PlayerDrawsCardsPayload", ws.PlayerDrawsCardsPayload{
		PlayerId: target.P.Id,
		Name:     target.P.Name,
		Amount:   amount,
	})
}
