package arbitrary

type Arbitrary[T any] interface {
	Generate() T
}
