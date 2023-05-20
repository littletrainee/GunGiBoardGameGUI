package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func (g *Game) SelectKoma() {
	// 循環查找哪個block是否有被按鍵釋放
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, currentBlock := range g.Board.Blocks {
			// 若當前block有被按鍵釋放
			if currentBlock.OnBlock(x, y) {
				if len(currentBlock.KomaStack) > 0 {
					tempblock := g.Board.Blocks[k]
					tempblock.BeSelect = true
					g.Board.Blocks[k] = tempblock
					g.WhichKomaBeenSelected = []int{k.X, k.Y}
					// 用暫時變數引用當前的block的最後一個駒
					lastKoma := currentBlock.KomaStack[len(currentBlock.KomaStack)-1]

					// 取得當前block的駒堆疊數量，以確認當前block最後一個駒可以移動的最大範圍
					currentDan := len(currentBlock.KomaStack)
					if lastKoma.Name == "帥" {
						if g.GameState.Level != level.INTERMEDIATE &&
							g.GameState.Level != level.ADVANCED {
							for _, direction := range lastKoma.ProbablyMoveing {
								for danIndex, eachDanCanMove := range direction {
									if danIndex < currentDan {
										g.ConfirmPosition = append(g.ConfirmPosition, suICheck(eachDanCanMove, lastKoma, g, currentDan)...)
									}
								}
							}
							g.Board.Blocks[k] = currentBlock
							g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
							return
						}
					}

					// 迭代當前的方向
					for _, direction := range lastKoma.ProbablyMoveing {
						g.ConfirmPosition = append(g.ConfirmPosition, checkTargetBlock(direction, lastKoma, currentDan, g)...)
					}
					g.Board.Blocks[k] = currentBlock
					g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
					return
				}
			}
		}

		// 選擇駒台的駒
		for i, v := range g.Player1.KomaTai {
			if v.Item1.OnKoma(float64(x), float64(y)) {
				if v.Item2 > 0 {
					g.WhichKomaBeenSelected = []int{i}
					// g.Player1.WhichOneSelected = i
					for k := range g.Board.Blocks {
						tempblock := g.Board.Blocks[k]
						tempblock.BeSelect = false
						tempblock.CurrentColor = color.BoardColor
						g.Board.Blocks[k] = tempblock
					}
					for column := 1; column < 10; column++ {
						for row := g.Player1.MaxRange; row < 10; row++ {
							tempblock := g.Board.Blocks[image.Point{X: column, Y: row}]
							if tempblock.HasSuI() {
								tempblock.CurrentColor = color.DenyColor
							} else if len(tempblock.KomaStack) < g.GameState.MaxLayer {
								tempblock.CurrentColor = color.ConfirmColor
							} else {
								tempblock.CurrentColor = color.DenyColor
							}
							g.Board.Blocks[image.Point{X: column, Y: row}] = tempblock
						}
					}
					g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
				}
			}
		}
	}
}
