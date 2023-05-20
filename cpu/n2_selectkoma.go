package cpu

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
)

func (c *CPU) SelectKoma(b board.Board) {
	if c.InevitableWinOpportunity(b) {
		return
	}
	if c.InevitableLoseOpportunity(b) {
		// try capture or control
		c.TryCapture(b)
		if len(c.TargetMove) > 0 {
			fmt.Println("GG Been Capture")
			return
		}
	}
	// BestMoveOption()
	// tryCapture()
	c.TryCapture(b)
	if len(c.TargetMove) > 0 {
		fmt.Println("Capture")
		return
	}

	var probablyChoice [][]int
	rand.Seed(time.Now().UnixNano())
	// 迭代block
	for k, v := range b.Blocks {
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color == c.Player.SelfColor {
			probablyChoice = append(probablyChoice, []int{k.X, k.Y})
		}
	}

	// 迭代駒台
	for i, v := range c.Player.KomaTai {
		if v.Item2 > 0 {
			probablyChoice = append(probablyChoice, []int{i})
		}
	}
	// 從可選擇列表選擇一個
	c.targetKoma = probablyChoice[rand.Intn(len(probablyChoice))]
}
