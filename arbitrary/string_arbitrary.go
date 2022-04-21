package arbitrary

import "strings"

type StringArbitrary struct {
	size int
}

func NewStringArbitrary(size int) *StringArbitrary {
	return &StringArbitrary{size}
}

func (g *StringArbitrary) Generate() string {
	charGenerator := NewCharArbitrary()
	stringBuilder := strings.Builder{}
	for i := 0; i < g.size; i++ {
		randomChar := charGenerator.Generate()
		stringBuilder.WriteRune(randomChar)
	}

	return stringBuilder.String()
}
