package uno

import (
	"math/rand"

	"uno_online/api/ws"
	"uno_online/game"
)

func reversePlayed(gp *game.GamePlayer, card *game.Card, state *game.GameState) {
	u, _ := (*card).(UnoCard)
	if u.Value != Reverse {
		return
	}
	state.CurrDir *= -1
	state.WsRoom.BroadcastMessage("DirectionChangedPayload", ws.DirectionChangedPayload{
		Direction: state.CurrDir,
	})
}

func skipPlayed(gp *game.GamePlayer, card *game.Card, state *game.GameState) {
	u, _ := (*card).(UnoCard)
	if u.Value != Skip {
		return
	}
	state.WsRoom.BroadcastMessage("PlayerSkippedPayload", ws.PlayerSkippedPayload{})
	state.NextPlayer()
}

func plusTwoPlayed(gp *game.GamePlayer, card *game.Card, state *game.GameState) {
	u, _ := (*card).(UnoCard)
	if u.Value != Plus2 {
		return
	}

	p := state.PeekNextPlayer()
	state.DrawCards(p, 2)
	state.NextPlayer()
	state.WsRoom.BroadcastMessage("PlayerSkippedPayload", ws.PlayerSkippedPayload{})
}

func wildcardPlayed(gp *game.GamePlayer, card *game.Card, state *game.GameState) {
	u, _ := (*card).(*UnoCard)
	if u.Color != Black || u.Value != Wildcard {
		return
	}

	choiceInt, timeout := game.AskColor(gp, colorsToInts(basicColors))
	choice := Color(choiceInt)
	if timeout {
		choice = basicColors[rand.Intn(len(basicColors))]
	}
	u.Chosen = choice

	state.WsRoom.BroadcastMessage("PlayerChoseColorPayload", ws.PlayerChoseColorPayload{
		PlayerId: gp.P.Id,
		Name:     gp.P.Name,
		Color:    choiceInt,
	})
	state.NextPlayer()
	state.WsRoom.BroadcastMessage("PlayerSkippedPayload", ws.PlayerSkippedPayload{})
}

func wildcard4Played(gp *game.GamePlayer, card *game.Card, state *game.GameState) {
	u, _ := (*card).(*UnoCard)
	if u.Color != Black || u.Value != Wildcard4 {
		return
	}

	choiceInt, timeout := game.AskColor(gp, colorsToInts(basicColors))
	choice := Color(choiceInt)
	if timeout {
		choice = basicColors[rand.Intn(len(basicColors))]
	}
	u.Chosen = choice

	state.WsRoom.BroadcastMessage("PlayerChoseColorPayload", ws.PlayerChoseColorPayload{
		PlayerId: gp.P.Id,
		Name:     gp.P.Name,
		Color:    choiceInt,
	})

	p := state.PeekNextPlayer()
	state.DrawCards(p, 4)

	state.NextPlayer()
	state.WsRoom.BroadcastMessage("PlayerSkippedPayload", ws.PlayerSkippedPayload{})
}

func UnoCardPlacedListeners() []game.CardPlayEventListener {
	return []game.CardPlayEventListener{reversePlayed, skipPlayed, plusTwoPlayed, wildcardPlayed, wildcard4Played}
}

func colorsToInts(colors []Color) []int {
	asInt := make([]int, len(colors))
	for i, c := range colors {
		asInt[i] = int(c)
	}
	return asInt
}
