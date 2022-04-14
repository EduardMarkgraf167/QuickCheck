package generators

import "math/rand"

type BoolGenerator struct {
}

func (g *BoolGenerator) Generate() bool {
	return rand.Int()%2 == 0
}

func NewBoolGenerator() *BoolGenerator {
	return &BoolGenerator{}
}
