package generators

import "math/rand"

type IntGenerator struct {
}

func (g *IntGenerator) Generate() int {
	return rand.Int()
}

func NewIntGenerator() *IntGenerator {
	return &IntGenerator{}
}
