package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func (g *Game) MoveOnBoard() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// 選取的駒是否為駒台上的駒
		if len(g.WhichKomaBeenSelected) == 1 {
			// 新
			for k, currentBlock := range g.Board.Blocks {
				if currentBlock.OnBlock(x, y) {
					// 複製駒台上的駒
					targetKoma := g.Player1.KomaTai[g.WhichKomaBeenSelected[0]].Item1.Clone()
					// 設置偏移量
					shift := block.Shift(currentBlock.KomaStack)
					// 對被複製的駒重設其位置與座標
					targetKoma.CurrentCoordinate.X = int(currentBlock.Coordinate.X+constant.BLOCK_SIZE/2) - shift
					targetKoma.CurrentCoordinate.Y = int(currentBlock.Coordinate.Y+constant.BLOCK_SIZE/2) - shift
					targetKoma.SetCurrentPosition(currentBlock.Name.Item1, currentBlock.Name.Item2)
					targetKoma.SetGeoMetry(0)

					// 當前block的KomaStack增加被複製的駒
					currentBlock.KomaStack = append(currentBlock.KomaStack, targetKoma)
					g.Player1.KomaTai[g.WhichKomaBeenSelected[0]].Item2--
					if g.Player1.KomaTai[g.WhichKomaBeenSelected[0]].Item2 == 0 {
						g.Player1.KomaTai[g.WhichKomaBeenSelected[0]].Item1 = koma.Koma{}
					}
					g.Board.Blocks[k] = currentBlock
					g.WhichKomaBeenSelected = nil

					for k := range g.Board.Blocks {
						tempblock := g.Board.Blocks[k]
						tempblock.BeSelect = false
						tempblock.CurrentColor = color.BoardColor
						tempblock.ConfirmPosition = nil
						g.Board.Blocks[k] = tempblock
					}
					g.delayedChangePhaseTo(phase.DUELING_PHASE_CLICK_CLOCK)
					g.SetOwnMaxRange()
					return
				}
			}
			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				tempblock.ConfirmPosition = nil
				g.Board.Blocks[k] = tempblock
			}
			g.WhichKomaBeenSelected = nil
			g.delayedChangePhaseTo(phase.DUELING_PHASE_SELECT_KOMA)

			// 非駒台上的駒
		} else {
			for k, currentBlock := range g.Board.Blocks {
				if currentBlock.OnBlock(x, y) {
					// targetposition := image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}
					// 若點選的位置為棋盤上被選擇駒的位置，則取消選取
					if k.X == g.WhichKomaBeenSelected[0] && k.Y == g.WhichKomaBeenSelected[1] {
						for k := range g.Board.Blocks {
							tempblock := g.Board.Blocks[k]
							tempblock.BeSelect = false
							tempblock.CurrentColor = color.BoardColor
							tempblock.ConfirmPosition = nil
							g.Board.Blocks[k] = tempblock
						}
						g.WhichKomaBeenSelected = nil
						g.delayedChangePhaseTo(phase.DUELING_PHASE_SELECT_KOMA)
						return
					}
					var containOpponentKoma bool
					for _, v := range currentBlock.KomaStack {
						if v.Color != g.Player1.SelfColor {
							containOpponentKoma = true
						}
					}
					// 若目標block的段數已經答這個階級所能堆疊到最高段數，並且有包含對家的駒
					if len(currentBlock.KomaStack) == g.GameState.MaxLayer && containOpponentKoma {
						g.Capture.Show = true
						g.Capture.CaptureBool = true
						g.delayedChangePhaseTo(phase.DUELING_PHASE_CAPTURE_OR_CONTROL_ASK)
						return
					}

					// 若目標block的段數未達這個階級所能堆疊的最高段數，並且有包含對家的駒
					if len(currentBlock.KomaStack) < g.GameState.MaxLayer && containOpponentKoma {
						g.Capture.Show = true
						g.Capture.CaptureBool = true
						g.Capture.ControlBool = true
						g.delayedChangePhaseTo(phase.DUELING_PHASE_CAPTURE_OR_CONTROL_ASK)
						// 若目標位置已經達最高段數且目標位置有包括對方的駒時
					} else {
						// 複製前一個block
						previousBlock := g.Board.Blocks[image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}]
						// 複製前一個block中KomaStack的最後一個駒
						targetKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
						shift := block.Shift(currentBlock.KomaStack)
						// 對被複製的駒重設其位置與座標
						targetKoma.CurrentCoordinate.X = int(currentBlock.Coordinate.X+constant.BLOCK_SIZE/2) - shift
						targetKoma.CurrentCoordinate.Y = int(currentBlock.Coordinate.Y+constant.BLOCK_SIZE/2) - shift
						targetKoma.SetCurrentPosition(currentBlock.Name.Item1, currentBlock.Name.Item2)
						targetKoma.SetGeoMetry(0)
						// 當前block的KomaStack增加被複製的駒
						currentBlock.KomaStack = append(currentBlock.KomaStack, targetKoma)
						// 前一個block中的KomaStack移除最後一個
						previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
						// 重設可移動的位置切片
						previousBlock.ConfirmPosition = nil

						// 將前一個block重新賦予回去map
						g.Board.Blocks[image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}] = previousBlock
						// 重設前一個被選取的位置
						g.WhichKomaBeenSelected = nil
						g.Board.Blocks[k] = currentBlock

						for k := range g.Board.Blocks {
							tempblock := g.Board.Blocks[k]
							tempblock.BeSelect = false
							tempblock.CurrentColor = color.BoardColor
							tempblock.ConfirmPosition = nil
							g.Board.Blocks[k] = tempblock
						}
						g.delayedChangePhaseTo(phase.DUELING_PHASE_CLICK_CLOCK)
						g.SetOwnMaxRange()
						return
					}
				}
			}
		}
	}
}
