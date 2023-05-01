package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

func (g *Game) MoveOnBoard() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		for k, currentBlock := range g.Board.Blocks {
			if currentBlock.OnBlock(x, y) {
				for _, v := range g.Board.Blocks[g.WhichBlockBeSelect].ConfirmPosition {
					if k == v {
						// 複製前一個block
						previousBlock := g.Board.Blocks[g.WhichBlockBeSelect]
						// 複製前一個block中KomaStack的最後一個駒
						targetKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
						shift := Shift(currentBlock.KomaStack)
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
						g.Board.Blocks[g.WhichBlockBeSelect] = previousBlock
						// 重設前一個被選取的位置
						g.WhichBlockBeSelect = image.Point{}
						g.Board.Blocks[k] = currentBlock

						for k := range g.Board.Blocks {
							tempblock := g.Board.Blocks[k]
							tempblock.BeSelect = false
							tempblock.CurrentColor = color.BoardColor
							tempblock.ConfirmPosition = nil
							g.Board.Blocks[k] = tempblock
						}
						g.delayedChangePhaseTo(phase.DUELING_PHASE_CLICK_CLOCK)
					}
				}

			}
		}
	}
}
