package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

// DuelingPhaseSelectKoma 對弈期間的選駒，可能是棋盤上也可能是玩家的駒台
func (g *Game) DuelingPhaseSelectKoma() {
	// 循環查找哪個block是否有被按鍵釋放
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, currentBlock := range g.Board.Blocks {
			// 若當前block有被按鍵釋放
			if currentBlock.OnBlock(x, y) {
				if len(currentBlock.KomaStack) > 0 && currentBlock.KomaStack[len(currentBlock.KomaStack)-1].Color == g.Player1.SelfColor {
					tempblock := g.Board.Blocks[k]
					g.Board.Blocks[k] = tempblock
					g.WhichKomaBeenSelected = []int{k.X, k.Y}
					// 用暫時變數引用當前的block的最後一個駒
					lastKoma := currentBlock.KomaStack[len(currentBlock.KomaStack)-1]

					// 取得當前block的駒堆疊數量，以確認當前block最後一個駒可以移動的最大範圍
					currentDan := len(currentBlock.KomaStack)
					switch lastKoma.Name {
					case "帥":
						if g.GameState.LevelHolder.CurrentLevel != level.INTERMEDIATE &&
							g.GameState.LevelHolder.CurrentLevel != level.ADVANCED {
							for _, direction := range lastKoma.MoveableRange {
								for danIndex, eachDanCanMove := range direction {
									if danIndex < currentDan {
										g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, suICheck(eachDanCanMove, lastKoma, g.Board, currentDan)...)
									}
								}
							}
							// g.Board.Blocks[k] = currentBlock
							g.delayedChangePhaseTo(phase.MOVE_KOMA)
							return
						}
					case "弓":
						for i, direction := range lastKoma.MoveableRange {
							switch i {
							case 0:
								temp := image.Point{X: k.X, Y: k.Y - 1}
								if len(g.Board.Blocks[temp].KomaStack) > len(currentBlock.KomaStack) {
									continue
								}
							case 1:
								temp := image.Point{X: k.X - 1, Y: k.Y - 1}
								if len(g.Board.Blocks[temp].KomaStack) > len(currentBlock.KomaStack) {
									continue
								}
							case 7:
								temp := image.Point{X: k.X + 1, Y: k.Y - 1}
								if len(g.Board.Blocks[temp].KomaStack) > len(currentBlock.KomaStack) {
									continue
								}
							}
							g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, otherCheck(direction, lastKoma, currentDan, g)...)
						}
						// g.Board.Blocks[k] = currentBlock
						g.delayedChangePhaseTo(phase.MOVE_KOMA)
						return
					case "砲":
						for i, direction := range lastKoma.MoveableRange {
							if i == 0 {
								temp1 := image.Point{X: k.X, Y: k.Y - 1}
								temp2 := image.Point{X: k.X, Y: k.Y - 2}
								if len(g.Board.Blocks[temp1].KomaStack) > len(currentBlock.KomaStack) ||
									len(g.Board.Blocks[temp2].KomaStack) > len(currentBlock.KomaStack) {
									continue
								}
							}
							g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, otherCheck(direction, lastKoma, currentDan, g)...)
						}
						// g.Board.Blocks[k] = currentBlock
						g.delayedChangePhaseTo(phase.MOVE_KOMA)
						return
					case "筒":
						for i, direction := range lastKoma.MoveableRange {
							if i == 0 {
								temp1 := image.Point{X: k.X, Y: k.Y - 1}
								if len(g.Board.Blocks[temp1].KomaStack) > len(currentBlock.KomaStack) {
									continue
								}
							}
							g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, otherCheck(direction, lastKoma, currentDan, g)...)
						}
						// g.Board.Blocks[k] = currentBlock
						g.delayedChangePhaseTo(phase.MOVE_KOMA)
						return
					default:
						// 迭代當前的方向
						for _, direction := range lastKoma.MoveableRange {
							g.ConfirmPositionSlice = append(g.ConfirmPositionSlice, otherCheck(direction, lastKoma, currentDan, g)...)
						}
						// g.Board.Blocks[k] = currentBlock
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
							if tempblock.HasSuI() {
								tempblock.CurrentColor = color.DenyColor
							} else if len(tempblock.KomaStack) < g.GameState.LevelHolder.MaxLayer {
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
