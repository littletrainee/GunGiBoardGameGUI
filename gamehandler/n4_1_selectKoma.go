package gamehandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func (g *Game) selectKoma() {
	// 循環查找哪個block是否有被按鍵釋放
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, currentBlock := range g.Board.Blocks {
		Recheck:
			// 若當前block有被按鍵釋放
			if currentBlock.OnBlock(x, y) {
				if len(currentBlock.KomaStack) > 0 && g.WhichBlockBeSelect == (image.Point{}) {
					tempblock := g.Board.Blocks[k]
					tempblock.BeSelect = true
					g.Board.Blocks[k] = tempblock
					g.WhichBlockBeSelect = k
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
										currentBlock.ConfirmPosition = append(
											currentBlock.ConfirmPosition,
											suICheck(eachDanCanMove, lastKoma, g, currentDan)...)
									}
								}
							}
							g.Board.Blocks[k] = currentBlock
							g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
							goto Out
						}
					}

					// 迭代當前的方向
					for _, direction := range lastKoma.ProbablyMoveing {
						for danIndex, eachDanCanMove := range direction {
							// 確定沒有超過當前階級可以的最高段數
							if danIndex < currentDan {
								currentBlock.ConfirmPosition = append(
									currentBlock.ConfirmPosition,
									checkTargetBlock(eachDanCanMove, lastKoma, g, currentDan)...)
							}
						}
					}
					g.Board.Blocks[k] = currentBlock
					g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
					goto Out
				} else {
					for k := range g.Board.Blocks {
						tempblock := g.Board.Blocks[k]
						tempblock.BeSelect = false
						tempblock.CurrentColor = color.BoardColor
						tempblock.ConfirmPosition = nil
						g.Board.Blocks[k] = tempblock
					}
					// 如果視同一個位置
					if k != g.WhichBlockBeSelect {
						g.WhichBlockBeSelect = image.Point{}
						g.Player1.KomaTaiBeingClick = false
						g.Player1.WhichOneSelected = -1
						goto Recheck
					}
					g.WhichBlockBeSelect = image.Point{}

					g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
					goto Out
				}
			}
		}

		for i, v := range g.Player1.KomaTai {
			if v.Item1.OnKoma(float64(x), float64(y)) {
				if v.Item2 > 0 && !g.Player1.KomaTaiBeingClick {
					g.Player1.KomaTaiBeingClick = true
					g.Player1.WhichOneSelected = i
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
				} else {
					g.Player1.KomaTaiBeingClick = false
					g.Player1.WhichOneSelected = -1
					// 延遲可選擇取消
					for k := range g.Board.Blocks {
						tempblock := g.Board.Blocks[k]
						tempblock.CurrentColor = color.BoardColor
						g.Board.Blocks[k] = tempblock
						g.delayedChangePhaseTo(phase.DUELING_PHASE_MOVE_KOMA)
					}
				}
			}
		}
	Out:
	}
}
