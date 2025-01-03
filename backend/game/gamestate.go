package game

import "uno_online/api/models"

type CardPlayEventListener func(gp *GamePlayer, card Card, state *GameState)

type GamePlayer struct {
	p    *models.Player
	Hand []Card
}

type GameState struct {
	Room    *models.Room
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

	return state.Players[state.CurrI]
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
		target.Hand = append(target.Hand, state.Deck.Draw())
	}
}
