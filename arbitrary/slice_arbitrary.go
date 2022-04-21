package arbitrary

type SliceArbitrary[T any] struct {
	arbitrary Arbitrary[T]
	size      int
}

func NewSliceArbitrary[T any](arbitrary Arbitrary[T], size int) *SliceArbitrary[T] {
	return &SliceArbitrary[T]{arbitrary, size}
}

func (g *SliceArbitrary[T]) Generate() []T {
	var result []T
	for i := 0; i < g.size; i++ {
		result = append(result, g.arbitrary.Generate())
	}

	return result
}
