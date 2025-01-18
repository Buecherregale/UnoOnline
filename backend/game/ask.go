package game

import (
	"slices"
	"time"
	"uno_online/api/ws"
)

func AskCard(player *GamePlayer) (*Card, bool) {
	cps := make([]interface{}, 0)
	for _, c := range player.Hand {
		var i interface{} = c
		cps = append(cps, i)
	}

	message := ws.AskCardPayload{
		Options: cps,
	}

	validCard := false
	var chosen Card

	for !validCard {
		reply, timeout, _ := player.WsP.AskAndWaitReply("AskCardPayload", message, time.Second*30)
		if timeout {
			return nil, true
		}
		p, success := ws.MsgToPayload[ws.AnswerCardPayload](*reply)
		if !success {
			player.WsP.SendError(403, "could not parse card")
			continue
		}
		card := p.Card.(Card)

		for _, handCard := range player.Hand {
			if handCard == card {
				chosen = card
				break
			}
		}
		player.WsP.SendError(403, "invalid card chosen")
	}

	return &chosen, false
}

func AskColor(player *GamePlayer, colors []int) (int, bool) {
	message := ws.AskColorPayload{
		Options: colors,
	}

	valid := false
	var chosen int

	for !valid {
		reply, timeout, _ := player.WsP.AskAndWaitReply("AskColorPayload", message, time.Second*30)
		if timeout {
			return -1, true
		}
		p, success := ws.MsgToPayload[ws.AnswerColorPayload](*reply)
		if !success {
			player.WsP.SendError(403, "could not parse color")
			continue
		}
		chosen = p.Chosen
		if slices.Contains(colors, chosen) {
			valid = true
		} else {
			player.WsP.SendError(403, "invalid color chosen")
		}
	}
	return chosen, false
}
