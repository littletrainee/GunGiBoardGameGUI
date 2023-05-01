package gamestate

import "image/color"

func (g *GameState) TrunToNext() color.Gray16 {
	if g.Turn == color.Black {
		g.Turn = color.White
	} else {
		g.Turn = color.Black
	}
	return g.Turn
}
