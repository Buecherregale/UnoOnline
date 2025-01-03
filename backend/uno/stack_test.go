package uno

import (
	"testing"
)

type SimpleCard struct {
	value int
}

func (c SimpleCard) matches(bot Card) bool {
	o, ok := bot.(SimpleCard)
	return ok && c.value == o.value
}

func TestStackLen(t *testing.T) {
	toTest := Stack{}

	if toTest.Len() != 0 {
		t.Errorf("expected len 0 for empty stack, but got%d\n", toTest.Len())
	}
	c1 := SimpleCard{value: 4}
	toTest.Play(c1)
	if toTest.Len() != 1 {
		t.Errorf("expected len 1 for stack, but got%d\n", toTest.Len())
	}
}

func TestPlay(t *testing.T) {
	toTest := Stack{}

	c1 := SimpleCard{value: 4}
	if !toTest.Play(c1) {
		t.Errorf("expected to play on empty stack\n")
	}
	c2 := SimpleCard{value: 4}
	if !toTest.Play(c2) {
		t.Errorf("expected to play with same value\n")
	}
	c3 := SimpleCard{value: 5}
	if toTest.Play(c3) {
		t.Errorf("expected to fail play with different value %d\n", c3.value)
	}
}

func TestGetTop(t *testing.T) {
	toTest := Stack{}

	top := toTest.GetTop()
	if top != nil {
		t.Errorf("expected nil for top of empty stack\n")
	}
	c1 := SimpleCard{value: 4}
	toTest.Play(c1)

	top = toTest.GetTop()
	if top == nil {
		t.Errorf("expected played card as top card\n")
	}
	c2 := SimpleCard{value: 4}
	toTest.Play(c2)

	top = toTest.GetTop()
	if top != c2 {
		t.Errorf("expected last played card as top card, but got %s\n", top)
	}
}

func TestRetrieveBottomCard(t *testing.T) {
	toTest := Stack{}

	if toTest.RetrieveBottomCard() != nil {
		t.Errorf("expected nil for bottom card of empty stack\n")
	}
	c1 := SimpleCard{value: 4}
	toTest.Play(c1)

	if toTest.RetrieveBottomCard() != nil {
		t.Errorf("expected nil for bottom card of len 1 stack\n")
	}
	c2 := SimpleCard{value: 4}
	toTest.Play(c2)

	if toTest.RetrieveBottomCard() != c1 {
		t.Errorf("expected first played card for bottom card of stack\n")
	}
	if toTest.Len() != 1 {
		t.Errorf("expected len 1 for length of stack after retrieving last card, but got%d\n", toTest.Len())
	}
}

func TestRetrieveBottomCards(t *testing.T) {
	toTest := Stack{}

	if toTest.RetrieveBottomCards() != nil {
		t.Errorf("expected nil for bottom cards of empty stack\n")
	}
	toTest.Play(SimpleCard{value: 4})

	if toTest.RetrieveBottomCards() != nil {
		t.Errorf("expected nil for bottom cards of len 1 stack\n")
	}
	toTest.Play(SimpleCard{value: 4})
	toTest.Play(SimpleCard{value: 4})

	ret := toTest.RetrieveBottomCards()
	if len(ret) != 2 {
		t.Errorf("expected 2 cards for bottom cards of stack, but got %d cards\n", len(ret))
	}
	if toTest.Len() != 1 {
		t.Errorf("expected 1 card in stack after retrieving 2, but got %d cards\n", toTest.Len())
	}
}
