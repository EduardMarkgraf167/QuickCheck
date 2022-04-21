package arbitrary

import "math/rand"

type PickRandomElementArbitrary[T any] struct {
	elements []T
	length   int
}

func NewPickRandomElementArbitrary[T any](elements []T) *PickRandomElementArbitrary[T] {
	length := len(elements)
	return &PickRandomElementArbitrary[T]{elements, length}
}

func (g *PickRandomElementArbitrary[T]) Generate() T {
	index := int(rand.Uint64()) % g.length
	return g.elements[index]
}
