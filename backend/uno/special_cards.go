package uno

import "uno_online/game"

func ReversePlayed(gp game.GamePlayer, card game.Card, state game.GameState) {
  state.CurrDir *= -1
}

func SkipPlayed(gp game.GamePlayer, card game.Card, state game.GameState) {
  state.NextPlayer()
}

func PlusTwoPlayed(gp game.GamePlayer, card game.Card, state game.GameState) {
  p := state.PeekNextPlayer()
  state.NextPlayer()
}

func WildcardPlayed(gp game.GamePlayer, card game.Card, state game.GameState) {
  p := state.PeekNextPlayer()
  var choice Color = Green

  state.NextPlayer()
}

