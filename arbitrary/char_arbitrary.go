package arbitrary

import "math/rand"

type CharArbitrary struct {
}

func (g *CharArbitrary) Generate() rune {
	runes := []rune("abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ß?!\"§$%&/()=*~'#_-.:,;<>|°^")
	return runes[rand.Intn(len(runes))]
}

func NewCharArbitrary() *CharArbitrary {
	return &CharArbitrary{}
}
