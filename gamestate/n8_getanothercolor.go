package gamestate

import "image/color"

func (g *GameState) TurnToNext() {
	if g.Turn == color.Black {
		g.Turn = color.White
	} else {
		g.Turn = color.Black
	}
}
