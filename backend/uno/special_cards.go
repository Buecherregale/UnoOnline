package uno

import "uno_online/game"

func reversePlayed(gp *game.GamePlayer, card game.Card, state *game.GameState) {
	state.CurrDir *= -1
}

func skipPlayed(gp *game.GamePlayer, card game.Card, state *game.GameState) {
	state.NextPlayer()
}

func plusTwoPlayed(gp *game.GamePlayer, card game.Card, state *game.GameState) {
	p := state.PeekNextPlayer()
	state.DrawCards(p, 2)
	state.NextPlayer()
}

func wildcardPlayed(gp *game.GamePlayer, card game.Card, state *game.GameState) {
	choice := game.PromptChoice(gp, []Color{Red, Green, Blue, Yellow})
	u, _ := card.(UnoCard)
	u.Chosen = choice

	state.NextPlayer()
}

func wildcard4Played(gp *game.GamePlayer, card game.Card, state *game.GameState) {
	choice := game.PromptChoice(gp, []Color{Red, Green, Blue, Yellow})
	u, _ := card.(UnoCard)
	u.Chosen = choice

	p := state.PeekNextPlayer()
	state.DrawCards(p, 4)

	state.NextPlayer()
}

func UnoCardPlacedListeners() []game.CardPlayEventListener {
	return []game.CardPlayEventListener{reversePlayed, skipPlayed, plusTwoPlayed, wildcardPlayed, wildcard4Played}
}
