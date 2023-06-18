package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// 俘獲或控制按鈕的判斷
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

					targetKoma.SetCurrentCoordinate(currentBlock.Coordinate, shift)
					targetKoma.SetCurrentPosition(currentBlock.Name)
					targetKoma.SetGeoMetry(0)
					tempKomaStack = append(tempKomaStack, targetKoma)

				}
			}
			targetKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
			shift := block.Shift(tempKomaStack)

			targetKoma.SetCurrentCoordinate(currentBlock.Coordinate, shift)
			targetKoma.SetCurrentPosition(currentBlock.Name)
			targetKoma.SetGeoMetry(0)

			tempKomaStack = append(tempKomaStack, targetKoma)
			previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
			currentBlock.KomaStack = tempKomaStack
			g.Board.Blocks[g.TargetPosition] = currentBlock
			g.Board.Blocks[previousPosition] = previousBlock
			g.ConfirmPosition = nil

			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				g.Board.Blocks[k] = tempblock
			}
			g.delayedChangePhaseTo(phase.CLICK_CLOCK)
			g.SetMaxRange()
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
			targetKoma.SetCurrentCoordinate(currentBlock.Coordinate, shift)
			targetKoma.SetCurrentPosition(currentBlock.Name)
			targetKoma.SetGeoMetry(0)
			// 當前block的KomaStack增加被複製的駒
			currentBlock.KomaStack = append(currentBlock.KomaStack, targetKoma)
			// 前一個block中的KomaStack移除最後一個
			previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
			// 重設可移動的位置切片
			g.ConfirmPosition = nil

			// 將前一個block重新賦予回去map
			g.Board.Blocks[image.Point{X: g.WhichKomaBeenSelected[0], Y: g.WhichKomaBeenSelected[1]}] = previousBlock
			// 重設前一個被選取的位置
			g.WhichKomaBeenSelected = nil
			g.Board.Blocks[g.TargetPosition] = currentBlock

			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				g.Board.Blocks[k] = tempblock
			}
			g.delayedChangePhaseTo(phase.CLICK_CLOCK)
			g.SetMaxRange()
			g.Capture.Reset()
			g.WhichKomaBeenSelected = nil
			g.TargetPosition = image.Point{}
		}
		if g.Capture.CancelButton(x, y) {
			for k := range g.Board.Blocks {
				tempblock := g.Board.Blocks[k]
				tempblock.BeSelect = false
				tempblock.CurrentColor = color.BoardColor
				g.Board.Blocks[k] = tempblock
			}
			g.delayedChangePhaseTo(phase.SELECT_KOMA)
			g.Capture.Reset()
			g.TargetPosition = image.Point{}
		}
	}
}
