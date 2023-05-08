package cpu

import (
	"image"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func (c *CPU) SelectMove(g gamestate.GameState, b board.Board) {
	var (
		currentLastKoma  koma.Koma
		probablyPosition []image.Point
		wall             bool
	)

	if len(c.targetKoma) == 1 { // 駒台
		for column := 1; column < 10; column++ {
			for row := 1; row < c.Player.MaxRange+1; row++ {
				targetPosition := image.Point{X: column, Y: row}
				targetBlock := b.Blocks[targetPosition]
				if len(targetBlock.KomaStack) < 2 && !targetBlock.HasSuI() {
					probablyPosition = append(probablyPosition, image.Point{X: column, Y: row})
				}
			}
		}

	} else { // 棋盤
		currentBlock := b.Blocks[image.Point{X: c.targetKoma[0], Y: c.targetKoma[1]}]
		currentLastKoma = currentBlock.KomaStack[len(currentBlock.KomaStack)-1]
		currentDan := len(currentBlock.KomaStack)
		if currentLastKoma.Name == "帥" {
			if g.Level != level.INTERMEDIATE && g.Level != level.ADVANCED {
				for _, direction := range currentLastKoma.ProbablyMoveing {
					for danIndex, eachDanCanMove := range direction {
						if danIndex < currentDan {
							probablyPosition = append(probablyPosition, suICheck(eachDanCanMove, currentLastKoma, b, currentDan)...)
						}
					}
				}
			}
		} else if currentLastKoma.Name == "大" {
			for d, v := range currentLastKoma.ProbablyMoveing {
				if d == 0 || d == 2 || d == 4 || d == 6 { // 方向
					wall = false
					for i := 0; i < len(v[0]); i++ {
						targetblockPosition := image.Point{
							X: currentLastKoma.CurrentPosition.Item1 - int(v[0][i].X),
							Y: currentLastKoma.CurrentPosition.Item2 + int(v[0][i].Y),
						}
						_, exists := b.Blocks[targetblockPosition]
						if exists {
							if !wall {
								targetBlock := b.Blocks[targetblockPosition]
								targetBlockLength := len(targetBlock.KomaStack)
								if targetBlockLength != 0 || targetBlockLength >= currentDan {
									wall = true
								}
								if targetBlockLength <= currentDan && targetBlockLength < g.MaxLayer && !targetBlock.HasSuI() {
									probablyPosition = append(probablyPosition, targetblockPosition)
								} else {
									break
								}
							} else {
								break
							}
						} else {
							break
						}
					}
				} else {
					for danIndex, eachDanCanMove := range v {
						if danIndex < currentDan {
							probablyPosition = append(probablyPosition, checkTargetBlock(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
						}
					}
				}
			}
		} else if currentLastKoma.Name == "中" {
			for d, v := range currentLastKoma.ProbablyMoveing {
				if d == 1 || d == 3 || d == 5 || d == 7 { // 方向
					wall = false
					for i := 0; i < len(v[0]); i++ {
						targetblockPosition := image.Point{
							X: currentLastKoma.CurrentPosition.Item1 - int(v[0][i].X),
							Y: currentLastKoma.CurrentPosition.Item2 + int(v[0][i].Y),
						}
						_, exists := b.Blocks[targetblockPosition]
						if exists {
							if !wall {
								targetBlock := b.Blocks[targetblockPosition]
								targetBlockLength := len(targetBlock.KomaStack)
								if targetBlockLength != 0 || targetBlockLength >= currentDan {
									wall = true
								}
								if targetBlockLength <= currentDan && targetBlockLength < g.MaxLayer && !targetBlock.HasSuI() {
									probablyPosition = append(probablyPosition, targetblockPosition)
								} else {
									break
								}
							} else {
								break
							}
						} else {
							break
						}
					}
				} else {
					for danIndex, eachDanCanMove := range v {
						if danIndex < currentDan {
							probablyPosition = append(probablyPosition, checkTargetBlock(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
						}
					}
				}
			}
		} else {
			for _, direction := range currentLastKoma.ProbablyMoveing {
				for danIndex, eachDanCanMove := range direction {
					// 確定沒有超過當前階級可以的最高段數
					if danIndex < currentDan {
						probablyPosition = append(probablyPosition,
							checkTargetBlock(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
					}
				}
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	c.targetPosition = probablyPosition[rand.Intn(len(probablyPosition))]

}
