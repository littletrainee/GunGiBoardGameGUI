package cpu

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (c *CPU) SelectKoma(b board.Board, gameState gamestate.GameState) {
	if c.inevitableWinOpportunity(b) {
		return
	}
	if c.inevitableLoseOpportunity(b) {
		// 嘗試俘獲防禦
		c.defenseCapture(b)
		if len(c.CaptureForDefense) > 0 {
			fmt.Println(currentTime(), ": can capture defense")
			return
		}
		fmt.Println(currentTime(), ": can't capture defense")

		// 嘗試移動防禦
		c.defenseAvoid(b)
		if len(c.AvoidForDefense) > 0 {
			fmt.Println(currentTime(), ": can move sui to avoid capture")
			return
		}
		fmt.Println(currentTime(), ": can't move sui to avoid capture")

		// try arata
		c.defenseARata(b, gameState.MaxLayer)
		if len(c.ARaTaForDefense) > 0 {
			fmt.Println(currentTime(), ": can arata koma to defense capture sui")
			return
		}
		fmt.Println(currentTime(), ": can't arata koma to defense capture sui")
		fmt.Println(currentTime(), "GG Been Capture")

	}

	// 嘗試俘獲
	c.tryCapture(b)
	if len(c.CaptureForDefense) > 0 {
		fmt.Println(currentTime(), ": try to Capture")
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
	for i, v := range c.KomaTai {
		if v.Item2 > 0 {
			probablyChoice = append(probablyChoice, []int{i})
		}
	}
	// 從可選擇列表選擇一個
	c.targetKoma = probablyChoice[rand.Intn(len(probablyChoice))]
}

func currentTime() string {
	return time.Now().Format("03:04:05 PM")
}
