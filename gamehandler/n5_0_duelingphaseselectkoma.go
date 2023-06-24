package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// DuelingPhaseSelectKoma 對弈期間的選駒，可能是棋盤上也可能是玩家的駒台
func (g *Game) DuelingPhaseSelectKoma() {
	// 循環查找哪個block是否有被按鍵釋放
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, currentBlock := range g.Board.Blocks {
			// 若當前block有被按鍵釋放
			if currentBlock.OnBlock(x, y) {
				// 目標位置有駒，最上方的駒是幾方的顏色
				if len(currentBlock.KomaStack) > 0 && currentBlock.KomaStack[len(currentBlock.KomaStack)-1].Color == g.Player1.SelfColor {
					// 設定選取的位置
					g.WhichKomaBeenSelected = []int{k.X, k.Y}
					switch currentBlock.KomaStack[len(currentBlock.KomaStack)-1].Name {
					case "帥":
						g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, suICheck(g.GameState.LevelHolder, currentBlock.KomaStack, g.Board)...)
						g.delayedChangePhaseTo(phase.MOVE_KOMA)
						return
					case "弓", "砲", "筒":
						g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, canHopCheck(g.GameState.LevelHolder, currentBlock.KomaStack, g.Board)...)
						g.delayedChangePhaseTo(phase.MOVE_KOMA)
						return
					default:
						g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, otherCheck(g.GameState.LevelHolder, currentBlock.KomaStack, g.Board)...)
						g.delayedChangePhaseTo(phase.MOVE_KOMA)
						return
					}

				}
			}
		}

		// 選擇駒台的駒
		for i, v := range g.Player1.KomaDai {
			if v.Item1.OnKoma(float64(x), float64(y)) {
				if v.Item2 > 0 {
					g.WhichKomaBeenSelected = []int{i}
					for k := range g.Board.Blocks {
						tempblock := g.Board.Blocks[k]
						tempblock.CurrentColor = color.BoardColor
						g.Board.Blocks[k] = tempblock
					}
					for column := 1; column < 10; column++ {
						for row := g.Player1.MaxRange; row < 10; row++ {
							tempblock := g.Board.Blocks[image.Point{X: column, Y: row}]
							tempblocklen := len(tempblock.KomaStack)

							// 目標位置有駒，沒有帥，段數未達到最高段數
							if tempblocklen > 0 && !tempblock.HasSuI() && g.GameState.LevelHolder.MaxLayer > tempblocklen || tempblocklen == 0 {
								tempblock.CurrentColor = color.ConfirmColor
							} else {
								tempblock.CurrentColor = color.DenyColor
							}
							g.Board.Blocks[image.Point{X: column, Y: row}] = tempblock
						}
					}
					g.delayedChangePhaseTo(phase.MOVE_KOMA)
				}
			}
		}
	}
}
