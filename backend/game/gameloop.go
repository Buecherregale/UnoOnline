package game

import (
	"uno_online/api/models"
	"uno_online/api/ws"
)

func StartRoom(room *models.Room, cards []Card, listeners []CardPlayEventListener) GameState {
	gps := make([]*GamePlayer, len(room.Players))
	for i, p := range room.Players {
		gps[i] = &GamePlayer{P: &p}
	}
	deck := NewDeck(cards)
	deck.Shuffle()

	state := GameState{
		Room:      room,
		WsRoom:    ws.Server.Rooms[room.Id],
		Players:   gps,
		Deck:      &deck,
		Stack:     &Stack{},
		Winner:    nil,
		CurrI:     0,
		CurrDir:   1,
		listeners: listeners,
	}

	state.deal(state.Players, 7)
	top := state.Deck.Draw()
	state.Stack.Play(top)
	state.WsRoom.BroadcastMessage("GameStartPayload", interface{}(top))

	return state
}

func (state *GameState) Run() {
	for state.Winner == nil {
		toMove := state.NextPlayer()

		if !toMove.canPlayACard(state.Stack.GetTop()) {
			state.DrawCards(toMove, 1)
		}
		if toMove.canPlayACard(state.Stack.GetTop()) {
			toMove.play(state)
			if toMove.checkWin() {
				state.Winner = toMove
			}
		}
	}
}

func (player *GamePlayer) checkWin() bool {
	return len(player.Hand) == 0
}

func (player *GamePlayer) play(state *GameState) {
	choice, timeout := AskCard(player)
	if timeout {
		return
	}
	played := state.Stack.Play(*choice)
	// skips in case of wrong choice (frontend can handle this)
	if played {
		state.WsRoom.BroadcastMessage("CardPlayedPayload", ws.CardPlayedPayload{
			PlayerId: player.P.Id,
			Name:     player.P.Name,
			Card:     interface{}(choice),
		})
		// notify listener
		for _, l := range state.listeners {
			l(player, choice, state)
		}
	}
}

func (player *GamePlayer) canPlayACard(top Card) bool {
	for _, c := range player.Hand {
		if c.Matches(top) {
			return true
		}
	}
	return false
}

func (state *GameState) deal(players []*GamePlayer, cardCount int) {
	for _, p := range players {
		hand := make([]Card, 0)
		p.Hand = hand
		state.DrawCards(p, cardCount)
	}
}
