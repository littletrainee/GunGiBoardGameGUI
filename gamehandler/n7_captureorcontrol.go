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

func (g *Game) CaptureOrControl() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if g.Capture.CaptureButton(x, y) {
			previousPosition := image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}
			previousBlock := g.Board.Blocks[previousPosition]
			tempKomaStack := []koma.Koma{}
			currentBlock := g.Board.Blocks[g.TargetPosition]
			for _, v := range currentBlock.KomaStack {
				if v.Color == g.Player1.SelfColor {
					targetKoma := v.Clone()
					currentBlock := g.Board.Blocks[g.TargetPosition]
					shift := block.Shift(tempKomaStack)

					targetKoma.CurrentCoordinate.X = int(currentBlock.Coordinate.X+constant.BLOCK_SIZE/2) - shift
					targetKoma.CurrentCoordinate.Y = int(currentBlock.Coordinate.Y+constant.BLOCK_SIZE/2) - shift
					targetKoma.SetCurrentPosition(currentBlock.Name.Item1, currentBlock.Name.Item2)
					targetKoma.SetGeoMetry(0)
					tempKomaStack = append(tempKomaStack, targetKoma)

				}
			}
			targetKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
			shift := block.Shift(tempKomaStack)

			targetKoma.CurrentCoordinate.X = int(currentBlock.Coordinate.X+constant.BLOCK_SIZE/2) - shift
			targetKoma.CurrentCoordinate.Y = int(currentBlock.Coordinate.Y+constant.BLOCK_SIZE/2) - shift
			targetKoma.SetCurrentPosition(currentBlock.Name.Item1, currentBlock.Name.Item2)
			targetKoma.SetGeoMetry(0)

			tempKomaStack = append(tempKomaStack, targetKoma)
			previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
			previousBlock.ConfirmPosition = nil
			currentBlock.KomaStack = tempKomaStack
			g.Board.Blocks[g.TargetPosition] = currentBlock
			g.Board.Blocks[previousPosition] = previousBlock

			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				tempblock.ConfirmPosition = nil
				g.Board.Blocks[k] = tempblock
			}
			g.delayedChangePhaseTo(phase.DUELING_PHASE_CLICK_CLOCK)
			g.SetOwnMaxRange()
			g.Capture.Reset()
			g.WhichKomaBeenSelected = nil
			g.TargetPosition = image.Point{}
		}

		if g.Capture.ControlButton(x, y) {
			// 複製前一個block
			previousBlock := g.Board.Blocks[image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}]
			// 複製前一個block中KomaStack的最後一個駒
			targetKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
			currentBlock := g.Board.Blocks[g.TargetPosition]
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
			g.Board.Blocks[g.TargetPosition] = currentBlock

			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				tempblock.ConfirmPosition = nil
				g.Board.Blocks[k] = tempblock
			}
			g.delayedChangePhaseTo(phase.DUELING_PHASE_CLICK_CLOCK)
			g.SetOwnMaxRange()
			g.Capture.Reset()
			g.WhichKomaBeenSelected = nil
			g.TargetPosition = image.Point{}
		}
		if g.Capture.CancelButton(x, y) {
			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				tempblock.ConfirmPosition = nil
				g.Board.Blocks[k] = tempblock
			}
			g.delayedChangePhaseTo(phase.DUELING_PHASE_SELECT_KOMA)
			g.Capture.Reset()
			g.TargetPosition = image.Point{}
		}
	}
}
