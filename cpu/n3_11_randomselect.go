package cpu

import (
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

// RandomSelect 隨機移動，不需要進行防禦的選擇，從可能的切片當中選擇要移動的位置，可能是新也可能是移居，取決於切片的長度
//
// 參數g為當前遊戲的狀態，b為棋盤物件
func (c *CPU) RandomSelect(g gamestate.GameState, b board.Board) {
	var (
		currentLastKoma  koma.Koma
		probablyPosition []image.Point
		hinder           bool
		probablyChoice   [][]int
	)
	rand.Seed(time.Now().UnixNano())
	// 迭代block
	for k, v := range b.Blocks {
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color == c.Player.SelfColor {
			probablyChoice = append(probablyChoice, []int{k.X, k.Y}) //駒在棋盤上的位置
		}
	}

	// 迭代駒台
	for i, v := range c.KomaDai {
		if v.Item2 > 0 {
			probablyChoice = append(probablyChoice, []int{i}) // 駒在駒台上的編號
		}
	}
	// 從可選擇列表選擇一個
	t := rand.Intn(len(probablyChoice)) // 隨機選擇在可能選擇的列表中的位置
	c.MoveToTarget = probablyChoice[t]  // 設定選擇的事可能選擇的列表中的位置
	c.Select = cpuselect.NORMAL
	fmt.Println(c.MoveToTarget)

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
		tempPos := image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
		currentBlock, ok := b.Blocks[tempPos]
		if ok {
			currentLastKoma = currentBlock.KomaStack[len(currentBlock.KomaStack)-1]
			currentDan := len(currentBlock.KomaStack)
			if currentLastKoma.Name == "帥" {
				if g.LevelHolder.CurrentLevel != level.INTERMEDIATE && g.LevelHolder.CurrentLevel != level.ADVANCED {
					probablyPosition = append(probablyPosition, checkSuIConfirmMove(b, tempPos)...)
				}
			} else if currentLastKoma.Name == "大" {
				for d, v := range currentLastKoma.MoveableRange {
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
									if targetBlockLength <= currentDan && targetBlockLength < g.LevelHolder.MaxLayer && !targetBlock.HasSuI() {
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
								probablyPosition = append(probablyPosition, otherCheck(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
							}
						}
					}
				}
			} else if currentLastKoma.Name == "中" {
				for d, v := range currentLastKoma.MoveableRange {
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
									if targetBlockLength <= currentDan && targetBlockLength < g.LevelHolder.MaxLayer && !targetBlock.HasSuI() {
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
								probablyPosition = append(probablyPosition, otherCheck(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
							}
						}
					}
				}
			} else if currentLastKoma.Name == "弓" {
				for i, direction := range currentLastKoma.MoveableRange {
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
							probablyPosition = append(probablyPosition, otherCheck(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
						}
					}

				}
			} else {
				for _, direction := range currentLastKoma.MoveableRange {
					for danIndex, eachDanCanMove := range direction {
						// 確定沒有超過當前階級可以的最高段數
						if danIndex < currentDan {
							probablyPosition = append(probablyPosition,
								otherCheck(eachDanCanMove, currentLastKoma, g, b, currentDan)...)
						}
					}
				}
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	if len(probablyPosition) == 0 {
		fmt.Println("check")
	}
	temp := probablyPosition[rand.Intn(len(probablyPosition))]
	c.MoveToTarget = append(c.MoveToTarget, temp.X, temp.Y)
	fmt.Println(c.MoveToTarget)
}
