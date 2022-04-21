package arbitrary

import "math/rand"

type BoolArbitrary struct {
}

func NewBoolArbitrary() *BoolArbitrary {
	return &BoolArbitrary{}
}

func (g *BoolArbitrary) Generate() bool {
	return rand.Int()%2 == 0
}
