package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// MutinyCheck 謀的叛變選擇視窗判斷
func (g *Game) MutinyCheck() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if g.Mutiny.MutinyButton(x, y) { // 執行背叛
			var (
				cloneKomaStack []koma.Koma
				targetBlock    block.Block = g.Board.Blocks[g.TargetPosition]
				targetBlockLen int         = len(targetBlock.KomaStack)
			)
			// 因為長度有targetBlockLen個，因此以這個為主去迭代
			for i := 0; i < targetBlockLen; i++ {
				// 若序號為最後一個駒則直接複製到目標位置後並中斷
				if i == targetBlockLen-1 {
					cloneKomaStack = append(cloneKomaStack, targetBlock.KomaStack[targetBlockLen-1])
					break
				}
				// 逐項檢查v[1]是否是targetBlock的KomaStack的序號
				for _, v := range g.Player1.MutinyList {
					// 如果是該序號則複製後增加到cloneKomaStack裡面，並中斷該迴圈
					if v[1] == i {
						// 從駒台拿出來，並設定
						cloneKoma := g.Player1.KomaDai[v[0]].Item1.Clone()
						cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, block.Shift(cloneKomaStack))
						cloneKoma.SetCurrentPosition(targetBlock.Name)
						cloneKoma.SetGeoMetry(0)

						cloneKomaStack = append(cloneKomaStack, cloneKoma)
						g.Player1.KomaDai[v[0]].Item2--
						break
					}
				}
			}
			targetBlock.KomaStack = cloneKomaStack
			g.Board.Blocks[g.TargetPosition] = targetBlock
			g.Player1.MutinyList = nil
			g.Player1.SelectBouShou = false
			g.TargetPosition = image.Point{}
			g.Mutiny.Reset()
			g.GameState.DelayedChangePhaseTo(phase.CLICK_CLOCK)
		} else if g.Mutiny.CancelButton(x, y) { // 執行取消
			g.Player1.MutinyList = nil
			g.Player1.SelectBouShou = false
			g.TargetPosition = image.Point{}
			g.Mutiny.Reset()
			g.GameState.DelayedChangePhaseTo(phase.CLICK_CLOCK)
		}
	}
}
