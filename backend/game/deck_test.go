package game

import (
	"testing"
)

func TestDeckLen(t *testing.T) {
	toTest := Deck{}

	if toTest.Len() != 0 {
		t.Errorf("expected len 0 for empty deck, but got%d\n", toTest.Len())
	}
	c1 := SimpleCard{value: 4}
	toTest.Add(c1)
	if toTest.Len() != 1 {
		t.Errorf("expected len 1 for deck, but got%d\n", toTest.Len())
	}
}

func TestDraw(t *testing.T) {
	toTest := Deck{}

	if toTest.Draw() != nil {
		t.Errorf("expected nil for draw card on empty deck\n")
	}

	c1 := SimpleCard{value: 4}
	toTest.Add(c1)
	if toTest.Draw() != c1 {
		t.Errorf("expected added card for draw on deck\n")
	}
	if toTest.Len() != 0 {
		t.Errorf("expected len 0 after drawing card from deck, but got %d\n", toTest.Len())
	}
}

func TestAdd(t *testing.T) {
	toTest := Deck{}

	c1 := SimpleCard{value: 4}
	toTest.Add(c1)
	toTest.Add(SimpleCard{value: 5})

	if toTest.Len() != 2 {
		t.Errorf("expected len 2 after adding 2 cards to deck, but got %d\n", toTest.Len())
	}
	if toTest.Draw() != c1 {
		t.Errorf("expected to draw first card added to deck, since adding should put the card to the bottom\n")
	}
}

func TestFill(t *testing.T) {
	toTest := Deck{}

	cards := []Card{SimpleCard{value: 3}, SimpleCard{value: 4}, SimpleCard{value: 5}}
	toTest.Fill(cards)
	if toTest.Len() != 3 {
		t.Errorf("expected len 3 after filling empty deck with 3 cards, but got %d\n", toTest.Len())
	}

	c1 := SimpleCard{value: 1}
	cards2 := []Card{c1}
	toTest.Fill(cards2)
	if toTest.Len() != 4 {
		t.Errorf("expected len 4 after filling deck with additional card, but got %d\n", toTest.Len())
	}
	if toTest.Draw() == c1 {
		t.Errorf("expected to draw different card from deck, expected to fill to the bottom\n")
	}
}

func TestNewDeck(t *testing.T) {
	toTest := NewDeck([]Card{SimpleCard{value: 1}, SimpleCard{value: 2}})
	if toTest.Len() != 2 {
		t.Errorf("expected 2 cards in newly created deck, but got %d\n", toTest.Len())
	}
}
