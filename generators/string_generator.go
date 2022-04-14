package generators

import "strings"

type StringGenerator struct {
	size int
}

func (g *StringGenerator) Generate() string {
	charGenerator := NewCharGenerator()
	stringBuilder := strings.Builder{}
	for i := 0; i < g.size; i++ {
		randomChar := charGenerator.Generate()
		stringBuilder.WriteRune(randomChar)
	}

	return stringBuilder.String()
}

func NewStringGenerator(size int) *StringGenerator {
	return &StringGenerator{size}
}
