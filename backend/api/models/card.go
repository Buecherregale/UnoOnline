package models

type Card struct {
  color Color 
  value Value 
}

type Color int 
type Value int 

const (
  red Color = iota 
)

const (
  one Value = iota
)
