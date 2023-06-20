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

func (c *CPU) DuelingPhaseSelectMove(g gamestate.GameState, b board.Board) {
	var (
		currentLastKoma  koma.Koma
		probablyPosition []image.Point
		hinder           bool
	)

	if len(c.MoveToTarget) == 1 { // 駒台
		for column := 1; column < 10; column++ {
			for row := 1; row < c.Player.MaxRange+1; row++ {
				targetPosition := image.Point{X: column, Y: row}
				targetBlock := b.Blocks[targetPosition]
				if len(targetBlock.KomaStack) < 2 && !targetBlock.HasSuI() {
					probablyPosition = append(probablyPosition, image.Point{X: column, Y: row})
				}
			}
		}

	} else if len(c.MoveToTarget) == 2 { // 棋盤
		currentBlock, ok := b.Blocks[image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}]
		if ok {
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
						hinder = false
						for i := 0; i < len(v[0]); i++ {
							targetblockPosition := image.Point{
								X: currentLastKoma.CurrentPosition.X - v[0][i].X,
								Y: currentLastKoma.CurrentPosition.Y + v[0][i].Y,
							}
							_, exists := b.Blocks[targetblockPosition]
							if exists {
								if !hinder {
									targetBlock := b.Blocks[targetblockPosition]
									targetBlockLength := len(targetBlock.KomaStack)
									if targetBlockLength != 0 || targetBlockLength >= currentDan {
										hinder = true
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
						hinder = false
						for i := 0; i < len(v[0]); i++ {
							targetblockPosition := image.Point{
								X: currentLastKoma.CurrentPosition.X - v[0][i].X,
								Y: currentLastKoma.CurrentPosition.Y + v[0][i].Y,
							}
							_, exists := b.Blocks[targetblockPosition]
							if exists {
								if !hinder {
									targetBlock := b.Blocks[targetblockPosition]
									targetBlockLength := len(targetBlock.KomaStack)
									if targetBlockLength != 0 || targetBlockLength >= currentDan {
										hinder = true
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
			} else if currentLastKoma.Name == "弓" {
				for i, direction := range currentLastKoma.ProbablyMoveing {
					switch i {
					case 0:
						if len(b.Blocks[image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1] + 1}].KomaStack) > len(currentBlock.KomaStack) {
							continue
						}
					case 1:
						if len(b.Blocks[image.Point{X: c.MoveToTarget[0] - 1, Y: c.MoveToTarget[1] + 1}].KomaStack) > len(currentBlock.KomaStack) {
							continue
						}
					case 7:
						if len(b.Blocks[image.Point{X: c.MoveToTarget[0] + 1, Y: c.MoveToTarget[1] + 1}].KomaStack) > len(currentBlock.KomaStack) {
							continue
						}
					}
					for j, eachDanCanMove := range direction {
						if j < currentDan {
							probablyPosition = append(probablyPosition, checkTargetBlock(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
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
	}
	rand.Seed(time.Now().UnixNano())
	temp := probablyPosition[rand.Intn(len(probablyPosition))]
	c.MoveToTarget = append(c.MoveToTarget, temp.X, temp.Y)

}
