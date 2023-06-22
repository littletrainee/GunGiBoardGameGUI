package gamestate

import "image/color"

func (g *GameState) TurnToNext() {
	if g.ColorHolder.Turn == color.Black {
		g.ColorHolder.Turn = color.White
	} else {
		g.ColorHolder.Turn = color.Black
	}
}
