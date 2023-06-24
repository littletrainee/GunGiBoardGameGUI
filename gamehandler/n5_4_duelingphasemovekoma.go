package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// DuelingPhaseMoveKoma 對弈期間的移駒，玩家選擇是否要取消選取、新、或是移駒
func (g *Game) DuelingPhaseMoveKoma() {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// 選取的駒是否為駒台上的駒
		if len(g.WhichKomaBeenSelected) == 1 {
			// 新
			for k, currentBlock := range g.Board.Blocks {
				if currentBlock.OnBlock(x, y) {
					// 複製駒台上的駒
					targetKoma := g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item1.Clone()
					// 設置偏移量
					// 對被複製的駒重設其位置與座標
					targetKoma.SetCurrentCoordinate(currentBlock.Coordinate, block.Shift(currentBlock.KomaStack))
					targetKoma.SetCurrentPosition(currentBlock.Name)
					targetKoma.SetGeoMetry(0)

					// 當前block的KomaStack增加被複製的駒
					currentBlock.KomaStack = append(currentBlock.KomaStack, targetKoma)
					g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item2--
					if g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item2 == 0 {
						g.Player1.KomaDai[g.WhichKomaBeenSelected[0]].Item1 = koma.Koma{}
					}
					g.Board.Blocks[k] = currentBlock
					g.WhichKomaBeenSelected = nil

					resetBlockColor(&g.Board)
					g.delayedChangePhaseTo(phase.CLICK_CLOCK)
					g.SetMaxRange()
					return
				}
			}
			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.CurrentColor = color.BoardColor
				g.Board.Blocks[k] = tempblock
			}
			g.WhichKomaBeenSelected = nil
			g.delayedChangePhaseTo(phase.SELECT_KOMA)

			// 非駒台上的駒
		} else {
			for k, targetBlock := range g.Board.Blocks {
				if targetBlock.OnBlock(x, y) {
					// 若點選的位置為棋盤上被選擇駒的位置，則取消選取
					if k.X == g.WhichKomaBeenSelected[0] && k.Y == g.WhichKomaBeenSelected[1] {
						for k := range g.Board.Blocks {
							tempblock := g.Board.Blocks[k]
							tempblock.CurrentColor = color.BoardColor
							g.Board.Blocks[k] = tempblock
						}
						g.WhichKomaBeenSelected = nil
						g.delayedChangePhaseTo(phase.SELECT_KOMA)
						return
					}
					// 目標位置是在核可的移動範圍內
					if otherfunction.Contain(g.ConfirmPositionSlice, k) {
						targetBlockLength := len(targetBlock.KomaStack)

						// 若目標位置有駒，並且目標最上面的駒是帥，且顏色不等於自家的顏色
						if targetBlockLength > 0 && targetBlock.KomaStack[targetBlockLength-1].Name == "帥" &&
							targetBlock.KomaStack[targetBlockLength-1].Color != g.Player1.SelfColor {
							g.Player1Timer.StopCountDown <- true
							g.AnotherRoundOrEnd.Show = true
							g.delayedChangePhaseTo(phase.ANOTHER_ROUND_OR_END)
							return
						}

						var containOpponentKoma bool
						for _, v := range targetBlock.KomaStack {
							if v.Color != g.Player1.SelfColor {
								containOpponentKoma = true
							}
						}
						previousPosition := image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}
						// 若目標block的段數已經答這個階級所能堆疊到最高段數，並且有包含對家的駒
						if targetBlockLength == g.GameState.LevelHolder.MaxLayer && containOpponentKoma &&
							targetBlockLength == len(g.Board.Blocks[previousPosition].KomaStack) {
							g.Capture.Show = true
							g.TargetPosition = k
							g.delayedChangePhaseTo(phase.PLAYER_CAPTURE_OR_CONTROL_ASK)
							return
						}

						// 若目標block的段數未達這個階級所能堆疊的最高段數，並且有包含對家的駒
						if targetBlockLength < g.GameState.LevelHolder.MaxLayer && containOpponentKoma {
							g.Capture.Show = true
							g.Capture.ControlBool = true
							g.TargetPosition = k
							g.delayedChangePhaseTo(phase.PLAYER_CAPTURE_OR_CONTROL_ASK)
							return
						}

						// 若目標位置已經達最高段數且目標位置有包括對方的駒時
						if targetBlockLength < g.GameState.LevelHolder.MaxLayer && !targetBlock.HasSuI() {
							g.delayedChangePhaseTo(phase.CLICK_CLOCK)
							// 複製前一個block
							previousBlock := g.Board.Blocks[previousPosition]
							// 複製前一個block中KomaStack的最後一個駒
							targetKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
							// 對被複製的駒重設其位置與座標
							targetKoma.SetCurrentCoordinate(targetBlock.Coordinate, block.Shift(targetBlock.KomaStack))
							targetKoma.SetCurrentPosition(targetBlock.Name)
							targetKoma.SetGeoMetry(0)
							// 當前block的KomaStack增加被複製的駒
							targetBlock.KomaStack = append(targetBlock.KomaStack, targetKoma)
							// 前一個block中的KomaStack移除最後一個
							previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
							// 重設可移動的位置切片
							g.ConfirmPositionSlice = nil

							// 將前一個block重新賦予回去map
							g.Board.Blocks[previousPosition] = previousBlock
							// 重設前一個被選取的位置
							g.WhichKomaBeenSelected = nil
							g.Board.Blocks[k] = targetBlock

							for k := range g.Board.Blocks {
								tempblock := g.Board.Blocks[k]
								tempblock.CurrentColor = color.BoardColor
								g.Board.Blocks[k] = tempblock
							}
							g.SetMaxRange()
							return
						}
					}
				}
			}
		}
	}
}
