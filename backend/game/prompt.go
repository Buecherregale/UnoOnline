package game

func PromptChoice[T any](target *GamePlayer, choices []T) T {
	return choices[0]
}
