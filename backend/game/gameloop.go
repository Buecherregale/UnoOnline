package game

import (
	"uno_online/api/models"
)

func StartRoom(room *models.Room, cards []Card, listeners []CardPlayEventListener) GameState {
	gps := make([]*GamePlayer, len(room.Players))
	for i, p := range room.Players {
		gps[i] = &GamePlayer{p: &p}
	}
	deck := NewDeck(cards)
	deck.Shuffle()

	state := GameState{
		Room:      room,
		Players:   gps,
		Deck:      &deck,
		Stack:     &Stack{},
		Winner:    nil,
		CurrI:     0,
		CurrDir:   1,
		listeners: listeners,
	}

	deal(state.Players, state.Deck, 7)
	state.Stack.Play(state.Deck.Draw())

	return state
}

func Run(state GameState) {
	for state.Winner == nil {
		//toMove := state.NextPlayer()

	}
}

func deal(players []*GamePlayer, deck *Deck, cardCount int) {
	for _, p := range players {
		hand := make([]Card, cardCount)
		for j := range cardCount {
			hand[j] = deck.Draw()
		}
		p.Hand = hand
	}
}
