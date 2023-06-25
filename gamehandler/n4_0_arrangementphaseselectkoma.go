package gamehandler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// ArrangementPhaseSelectKoma 布陣階段的選駒，玩家在布陣階段的選擇
func (g *Game) ArrangementPhaseSelectKoma() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if !g.Player1.IsSetSuI { // 非第一巡
			if g.Player1.KomaDai[0].Item1.OnKoma(float64(x), float64(y)) && g.Player1.KomaDai[0].Item2 > 0 {
				g.setConfirmPosition()
				g.WhichKomaBeenSelected = []int{0}
				g.Player1.IsSetSuI = true
				g.GameState.DelayedChangePhaseTo(phase.MOVE_KOMA)
			}
		} else { //第一巡
			for i := range g.Player1.KomaDai {
				if g.Player1.KomaDai[i].Item1.OnKoma(float64(x), float64(y)) {
					// 逐項迭代駒台上的位置
					g.setConfirmPosition()
					g.WhichKomaBeenSelected = []int{i}
					g.GameState.DelayedChangePhaseTo(phase.MOVE_KOMA)
				}
			}
		}
	}

}
