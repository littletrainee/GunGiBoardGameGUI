package gamestate

import (
	"image/color"
	"math/rand"
)

func (g *GameState) SetFirst() {
	target := rand.Intn(2)
	if target == 0 {
		g.First = color.Black
	} else {
		g.First = color.White
	}
}
