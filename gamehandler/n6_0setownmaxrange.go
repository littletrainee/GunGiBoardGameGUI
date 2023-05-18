package gamehandler

import "math"

func (g *Game) SetOwnMaxRange() {
	var fartest int = math.MaxInt64
	for k, v := range g.Board.Blocks {
		if len(v.KomaStack) != 0 && v.KomaStack[len(v.KomaStack)-1].Color == g.Player1.SelfColor {
			if k.Y < int(fartest) {
				fartest = k.Y
			}
		}
	}
	g.Player1.MaxRange = fartest
}
