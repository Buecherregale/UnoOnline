package uno

import "uno_online/models"

type CardPlayEventListener interface {
  OnCardPlayed(gp GamePlayer, card Card, state GameState)
}

type GamePlayer struct {
  p *models.Player
  Hand []Card
}

type GameState struct {
  Room      *models.Room
  Players   []*GamePlayer
  Deck      *Deck
  Stack     *Stack
  Winner    *GamePlayer

  currI     int    
  listeners []CardPlayEventListener
}

func (state *GameState) NextPlayer() {
  state.currI += 1
  state.currI %= len(state.Players)
}

func (state *GameState) RegisterListener(listener CardPlayEventListener) {
  state.listeners = append(state.listeners, listener)
}

func StartRoom(room *models.Room) GameState {
  gps := make([]*GamePlayer, len(room.Players))
  for i, p := range room.Players {
    gps[i] = &GamePlayer{p: &p}
  }
  deck := NewDeck(UnoCards())
  deck.Shuffle()
  
  state := GameState{
    Room: room,
    Players: gps,
    Deck: &deck,
    Stack: &Stack{},
    listeners: make([]CardPlayEventListener, 0),
  }

  deal(state.Players, state.Deck, 7)
  state.Stack.Play(state.Deck.Draw())

  return state
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

