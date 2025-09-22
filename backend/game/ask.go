package game

import (
	"slices"
	"time"
	"uno_online/api/ws"
)

// Returns: chosen card, valid card chosen, request did timeout
func AskCard(player *GamePlayer) (*Card, bool, bool) {
	cps := make([]any, 0)
	for _, c := range player.Hand {
		var i any = c
		cps = append(cps, i)
	}

	message := ws.AskCardPayload{
		Options: cps,
	}

	// use valid card to determine := draw/play
	validCard := false
	var chosen *Card

	for chosen != nil {
		reply, timeout, _ := player.WsP.AskAndWaitReply("AskCardPayload", message, time.Second*30)
		if timeout {
			return nil, false, true
		}
		p, success := ws.MsgToPayload[ws.AnswerCardPayload](*reply)
		if !success {
			player.WsP.SendError(403, "could not parse card")
			continue
		}
		card := p.Card.(Card)

		if slices.Contains(player.Hand, card) {
			chosen = &card
			validCard = true
			break;
		}
	}

	return chosen, validCard, false
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
