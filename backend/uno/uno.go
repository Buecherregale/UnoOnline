package uno

type UnoCard struct {
  Color Color 
  Value Value
  Chosen Color // for choosing colors e.g. wildcard 
}

func (card UnoCard) matches(bot Card) bool {
  o, ok := bot.(UnoCard)
  if !ok {
    return false
  }

  if card.Color == Black {
    return true 
  }

  if o.Color == Black {
    return o.Chosen == card.Color
  }

  return card.Color == o.Color || card.Value == o.Value
}

type Color int 
const (
  Red Color = iota
  Blue 
  Yellow
  Green
  Black
)
var basicColors = []Color{Red, Blue, Yellow, Green}
var specialColors = []Color{Black}

type Value int 
const (
  Zero Value = iota
  One 
  Two 
  Three
  Four
  Five
  Six
  Seven
  Eight
  Nine
  Skip
  Reverse
  Plus2
  Wildcard
  Wildcard4
)
var basicValues = []Value{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Skip, Reverse, Plus2}
var specialValues = []Value{Wildcard, Wildcard4}

func UnoCards() []Card {
  cards := make([]Card, 0)
  for _, c := range basicColors {
    for _, v := range basicValues {
      for range 4 {
        cards = append(cards, UnoCard{Color: c, Value: v})
      }
    }
  }
  for _, c := range specialColors {
    for _, v := range specialValues {
      for range 2 {
        cards = append(cards, UnoCard{Color: c, Value: v})
      }
    }
  }
  return cards
}
