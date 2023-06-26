package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// DuelingPhaseSelectKoma 對弈期間的選駒，可能是棋盤上也可能是玩家的駒台
func (g *Game) DuelingPhaseSelectKoma() {
	// 循環查找哪個block是否有被按鍵釋放
	// if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, currentBlock := range g.Board.Blocks {
			// 若當前block有被按鍵釋放
			if currentBlock.OnBlock(x, y) {
				// 目標位置有駒，最上方的駒是幾方的顏色
				if len(currentBlock.KomaStack) > 0 && currentBlock.KomaStack[len(currentBlock.KomaStack)-1].Color == g.Player1.SelfColor {
					// 設定選取的位置
					g.WhichKomaBeenSelected = []int{k.X, k.Y}
					currentKoma := currentBlock.KomaStack[len(currentBlock.KomaStack)-1]
					switch currentKoma.Name {
					case "帥":
						g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, suICheck(g.GameState.LevelHolder, currentBlock.KomaStack, g.Board)...)
						g.GameState.DelayedChangePhaseTo(phase.MOVE_KOMA)
						return
					case "弓", "砲", "筒":
						g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, canHopCheck(g.GameState.LevelHolder, currentBlock.KomaStack, g.Board)...)
						g.GameState.DelayedChangePhaseTo(phase.MOVE_KOMA)
						return
					default:
						if currentKoma.Name == "謀" {
							g.Player1.SelectBouShou = true
						}
						g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, otherCheck(g.GameState.LevelHolder, currentBlock.KomaStack, g.Board)...)
						g.GameState.DelayedChangePhaseTo(phase.MOVE_KOMA)
						return
					}

				}
			}
		}

		// 選擇駒台的駒
		for i, v := range g.Player1.KomaDai {
			if len(v) > 0 && v[0].OnKoma(float64(x), float64(y)) {
				g.WhichKomaBeenSelected = []int{i}
				for column := 1; column < 10; column++ {
					for row := g.Player1.MaxRange; row < 10; row++ {
						targetBlockPosition := image.Point{X: column, Y: row}
						targetBlock := g.Board.Blocks[targetBlockPosition]
						targetBlockLen := len(targetBlock.KomaStack)

						// 目標位置有駒，沒有帥，段數未達到最高段數
						if targetBlockLen > 0 && !targetBlock.HasSuI() &&
							g.GameState.LevelHolder.MaxLayer > targetBlockLen &&
							targetBlock.KomaStack[targetBlockLen-1].Color == g.Player1.SelfColor ||
							targetBlockLen == 0 {
							targetBlock.CurrentColor = color.ConfirmColor
							g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, targetBlockPosition)
						} else {
							targetBlock.CurrentColor = color.DenyColor
						}
						g.Board.Blocks[image.Point{X: column, Y: row}] = targetBlock
					}
				}
				g.GameState.DelayedChangePhaseTo(phase.MOVE_KOMA)
			}
		}
	}
}
