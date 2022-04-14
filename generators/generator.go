package generators

type Generator[T any] interface {
	Generate() T
}
