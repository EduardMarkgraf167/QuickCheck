package arbitrary

import "math/rand"

type IntArbitrary struct {
}

func NewIntArbitrary() *IntArbitrary {
	return &IntArbitrary{}
}

func (g *IntArbitrary) Generate() int {
	return rand.Int()
}
