package generators

import "math/rand"

type PickRandomElementGenerator[T any] struct {
	elements []T
	length   int
}

func NewPickRandomElementGenerator[T any](elements []T) *PickRandomElementGenerator[T] {
	length := len(elements)
	return &PickRandomElementGenerator[T]{elements, length}
}

func (g *PickRandomElementGenerator[T]) Generate() T {
	index := int(rand.Uint64()) % g.length
	return g.elements[index]
}
