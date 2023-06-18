package gamehandler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// 駒被選取時
func (g *Game) ArrangementPhaseMoveKoma() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, v := range g.Board.Blocks {
			if v.OnBlock(x, y) && !v.HasSuI() && v.Name.Item2 > 6 && len(v.KomaStack) < g.GameState.MaxLayer {
				//複製出來
				var tempKoma koma.Koma = g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item1.Clone()

				// 設定目標格的位置
				tempKoma.SetCurrentCoordinate(v.Coordinate, int(block.Shift(v.KomaStack)))
				//設定目標代號
				tempKoma.SetCurrentPosition(v.Name)
				// 設定文字顯示的位置
				tempKoma.SetGeoMetry(0)

				v.KomaStack = append(v.KomaStack, tempKoma)
				g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item2--
				if g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item2 == 0 {
					g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item1 = koma.Koma{}
				}
				g.Board.Blocks[k] = v
				ResetBlockColor(&g.Board)
				g.WhichKomaBeenSelected = nil
				g.GameState.Phase = phase.CLICK_CLOCK
			}
		}
	}
}
