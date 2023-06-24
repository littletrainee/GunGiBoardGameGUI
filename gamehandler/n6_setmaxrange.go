package gamehandler

import (
	"math"
)

// SetMaxRange 設定可以雙方可以選擇的最大範圍
func (g *Game) SetMaxRange() {
	var fartest int = math.MaxInt64
	for k, v := range g.Board.Blocks {
		if len(v.KomaStack) != 0 && v.KomaStack[len(v.KomaStack)-1].Color == g.Player1.SelfColor {
			if k.Y < fartest {
				fartest = k.Y
			}
		}
	}
	// 若超過7則最大7，也就是至少有靠自己的3列可以放置
	if fartest >= 7 {
		g.Player1.MaxRange = 7
	} else {
		g.Player1.MaxRange = fartest
	}
	fartest = 0
	for k, v := range g.Board.Blocks {
		if len(v.KomaStack) != 0 && v.KomaStack[len(v.KomaStack)-1].Color == g.Player2.SelfColor {
			if fartest < k.Y {
				fartest = k.Y
			}
		}
	}
	if fartest <= 3 {
		g.Player2.MaxRange = 3
	} else {
		g.Player2.MaxRange = fartest
	}
}
