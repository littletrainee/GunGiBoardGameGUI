package gamehandler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// ArrangementPhaseMoveKoma 布陣階段駒被選取時，玩家選擇放置的位置或是取消選擇
func (g *Game) ArrangementPhaseMoveKoma() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// 若點到棋盤上的位置則放置
		for k, v := range g.Board.Blocks {
			if v.OnBlock(x, y) && !v.HasSuI() && v.Name.Y > 6 &&
				len(v.KomaStack) < g.GameState.LevelHolder.MaxLayer {
				//複製出來
				var tempKoma koma.Koma = g.Player1.KomaDai[g.WhichKomaBeenSelected[0]][0].Clone()
				tempKomaLen := len(g.Player1.KomaDai[g.WhichKomaBeenSelected[0]])

				tempKoma.SetCurrentCoordinate(v.Coordinate, block.Shift(v.KomaStack))
				tempKoma.SetCurrentPosition(v.Name)
				tempKoma.SetGeoMetry(0)

				v.KomaStack = append(v.KomaStack, tempKoma)
				// g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item2--
				g.Player1.KomaDai[g.WhichKomaBeenSelected[0]] = g.Player1.KomaDai[g.WhichKomaBeenSelected[0]][:tempKomaLen-1]
				g.Board.Blocks[k] = v
				resetBlockColor(&g.Board)
				g.WhichKomaBeenSelected = nil
				g.GameState.DelayedChangePhaseTo(phase.CLICK_CLOCK)
				return
			}
		}
		// 弱點到駒台上被選取駒
		for i, v := range g.Player1.KomaDai {
			if v[0].OnKoma(float64(x), float64(y)) && i == g.WhichKomaBeenSelected[0] {
				resetBlockColor(&g.Board)
				g.WhichKomaBeenSelected = nil
				g.GameState.DelayedChangePhaseTo(phase.SELECT_KOMA)
				return
			}
		}
	}
}
