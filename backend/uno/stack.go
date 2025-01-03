package uno

type Stack struct {
  cards []Card
}

func (stack *Stack) GetTop() Card {
  if stack.Len() == 0 {
    return nil
  }
  return stack.cards[len(stack.cards) - 1]
}

func (stack *Stack) Play(c Card) bool { 
  if stack.Len() > 0 && !c.matches(stack.GetTop()) {
    return false
  }
  stack.cards = append(stack.cards, c)
  return true 
}

func (stack *Stack) RetrieveBottomCard() Card {
  if stack.Len() <= 1 {
    return nil 
  }
  c := stack.cards[0]
  stack.cards = stack.cards[1:]
  return c
}

func (stack *Stack) RetrieveBottomCards() []Card { 
  todo := stack.Len() - 1
  if todo <= 0 {
    return nil
  }
  cs := stack.cards[:todo]
  stack.cards = stack.cards[todo:]

  return cs
}

func (stack *Stack) Len() int {
  return len(stack.cards)
}
