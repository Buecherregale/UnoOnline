package uno

import (
	"math/rand"
)

type Card interface {
	matches(bot Card) bool
}

type Deck struct {
	cards []Card
}

func NewDeck(cards []Card) Deck {
	return Deck{cards: cards}
}

func (deck *Deck) Len() int {
	return len(deck.cards)
}

func (deck *Deck) Draw() Card {
	if deck.Len() == 0 {
		return nil
	}
	c := deck.cards[0]
	deck.cards = deck.cards[1:]

	return c
}

func (deck *Deck) Add(card Card) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) Fill(cards []Card) {
	deck.cards = append(deck.cards, cards...)
}

func (deck *Deck) Shuffle() {
	n := len(deck.cards)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}
