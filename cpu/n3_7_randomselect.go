package cpu

import (
	"image"
	"log"
	"math/rand"

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
	// 迭代block
	for k, v := range b.Blocks {
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color == c.Player.SelfColor {
			probablyChoice = append(probablyChoice, []int{k.X, k.Y}) //駒在棋盤上的位置
		}
	}

	// 迭代駒台
	for i, v := range c.KomaDai {
		if len(v) > 0 {
			probablyChoice = append(probablyChoice, []int{i}) // 駒在駒台上的編號
		}
	}

ReSelect:
	// 從可選擇列表選擇一個
	t := rand.Intn(len(probablyChoice)) // 隨機選擇在可能選擇的列表中的位置
	c.MoveToTarget = probablyChoice[t] // 設定選擇的事可能選擇的列表中的位置
	c.Select = cpuselect.RANDOM_SELECT

	// 選擇為駒台的駒
	if len(c.MoveToTarget) == 1 { // 駒台
		for column := 1; column < 10; column++ {
			for row := 1; row < c.Player.MaxRange+1; row++ {
				targetBoockPosition := image.Point{X: column, Y: row}
				targetBlock := b.Blocks[targetBoockPosition]
				targetBlockLen := len(targetBlock.KomaStack)

				// 目標位置有駒，沒有帥，段數未達到最高段樹
				if targetBlockLen > 0 && !targetBlock.HasSuI() &&
					g.LevelHolder.MaxLayer > targetBlockLen &&
					targetBlock.KomaStack[targetBlockLen-1].Color == c.SelfColor ||
					targetBlockLen == 0 {
					probablyPosition = append(probablyPosition, targetBoockPosition)
				}
			}
		}
		if len(probablyPosition) == 0 {
			log.Printf("選擇駒台上的駒階段probablyPosition是空的\n\n")
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
				log.Printf("%v\n原先的選擇是:%v，要重新選擇!\n\n", err, c.MoveToTarget)
				goto ReSelect
			}
			probablyPosition = append(probablyPosition, temp...)
		case "大", "中":
			temp, err := taiChuUCheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Printf("%v\n原先的選擇是:%v，要重新選擇!\n\n", err, c.MoveToTarget)
				goto ReSelect
			}
			probablyPosition = append(probablyPosition, temp...)
		case "弓", "砲", "筒":
			temp, err := canHopCheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Printf("%v\n原先的選擇是:%v，要重新選擇!\n\n", err, c.MoveToTarget)
				goto ReSelect
			}
			probablyPosition = append(probablyPosition, temp...)
		default:
			temp, err := otherCheck(g.LevelHolder, selectBlock.KomaStack, b)
			if err != nil {
				log.Printf("%v\n原先的選擇是:%v，要重新選擇!\n\n", err, c.MoveToTarget)
				goto ReSelect
			}
			probablyPosition = append(probablyPosition, temp...)
		}
	}
	if len(probablyPosition) == 0 {
		log.Println("probablyPosition是空的")
	}
	temp := probablyPosition[rand.Intn(len(probablyPosition))]
	c.MoveToTarget = append(c.MoveToTarget, temp.X, temp.Y)
}
