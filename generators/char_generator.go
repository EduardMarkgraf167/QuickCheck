package generators

import "math/rand"

type CharGenerator struct {
}

func (g *CharGenerator) Generate() rune {
	runes := []rune("abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ß?!\"§$%&/()=*~'#_-.:,;<>|°^")
	return runes[rand.Intn(len(runes))]
}

func NewCharGenerator() *CharGenerator {
	return &CharGenerator{}
}
