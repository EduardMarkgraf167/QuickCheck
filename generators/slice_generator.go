package generators

type SliceGenerator[T any] struct {
	generator Generator[T]
	size      int
}

func NewSliceGenerator[T any](generator Generator[T], size int) *SliceGenerator[T] {
	return &SliceGenerator[T]{generator, size}
}

func (g *SliceGenerator[T]) Generate() []T {
	var result []T
	for i := 0; i < g.size; i++ {
		result = append(result, g.generator.Generate())
	}

	return result
}
