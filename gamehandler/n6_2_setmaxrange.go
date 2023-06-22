package gamehandler

import "math"

// SetMaxRange 設定可以選擇的最大範圍
func (g *Game) SetMaxRange() {
	var fartest int = math.MaxInt64
	for k, v := range g.Board.Blocks {
		if len(v.KomaStack) != 0 && v.KomaStack[len(v.KomaStack)-1].Color == g.Player1.SelfColor {
			if k.Y < int(fartest) {
				fartest = k.Y
			}
		}
	}
	if fartest >= 7 {
		g.Player1.MaxRange = 7
	} else {
		g.Player1.MaxRange = fartest
	}
}
