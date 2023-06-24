package cpu

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

// RandomSelect 隨機移動，不需要進行防禦的選擇，從可能的切片當中選擇要移動的位置，可能是新也可能是移居，取決於切片的長度
//
// 參數g為當前遊戲的狀態，b為棋盤物件
func (c *CPU) RandomSelect(g gamestate.GameState, b board.Board) {
	var (
		probablyPosition []image.Point
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
	c.Select = cpuselect.RANDOM_SELECT

	// 選擇為駒台的駒
	if len(c.MoveToTarget) == 1 { // 駒台
		for column := 1; column < 10; column++ {
			for row := 1; row < c.Player.MaxRange+1; row++ {
				targetPosition := image.Point{X: column, Y: row}
				targetBlock := b.Blocks[targetPosition]
				targetBlockLen := len(targetBlock.KomaStack)

				// 目標位置有駒，沒有帥，段數未達到最高段樹
				if targetBlockLen < g.LevelHolder.MaxLayer && !targetBlock.HasSuI() {
					probablyPosition = append(probablyPosition, image.Point{X: column, Y: row})
				}
			}
		}
		if len(probablyPosition) == 0 {
			log.Println("選擇駒台上的駒階段probablyPosition是空的")
		}

	} else if len(c.MoveToTarget) == 2 { // 棋盤
		// 設定目標位置
		selectPosition := image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
		// 從blocks取出目標block
		selectBlock := b.Blocks[selectPosition]

		switch selectBlock.KomaStack[len(selectBlock.KomaStack)-1].Name {
		case "帥":
			temp, err := suICheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Println(err)
			}
			probablyPosition = append(probablyPosition, temp...)
		case "大", "中":
			temp, err := taiChuUCheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Println(err)
			}
			probablyPosition = append(probablyPosition, temp...)
		case "弓", "砲", "筒":
			temp, err := canHopCheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Println(err)
			}
			probablyPosition = append(probablyPosition, temp...)
		default:
			temp, err := otherCheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Println(err)
			}
			probablyPosition = append(probablyPosition, temp...)
		}
	}
	rand.Seed(time.Now().UnixNano())
	if len(probablyPosition) == 0 {
		log.Println("probablyPosition是空的")
	}
	temp := probablyPosition[rand.Intn(len(probablyPosition))]
	c.MoveToTarget = append(c.MoveToTarget, temp.X, temp.Y)
	fmt.Println(c.MoveToTarget)
}
