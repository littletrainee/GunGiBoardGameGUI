package cpu

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/anotherroundorend"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (c *CPU) DuelingPhaseSelectKoma(b board.Board, gameState gamestate.GameState, anotherRoundOrEnd *anotherroundorend.AnotherRoundOrEnd) {
	if c.inevitableWinOpportunity(b) {
		return
	}
	if c.inevitableLoseOpportunity(b) {
		// 嘗試俘獲防禦
		c.defenseCapture(b)
		if len(c.MoveToTarget) > 0 {
			fmt.Println(currentTime(), ": can capture defense")
			c.Select = cpuselect.DEFENSE_CAPTURE
			return
		}
		fmt.Println(currentTime(), ": can't capture defense")

		// 嘗試移動防禦
		c.defenseAvoid(b)
		if len(c.MoveToTarget) > 0 {
			fmt.Println(currentTime(), ": can move sui to avoid capture")
			c.Select = cpuselect.DEFENSE_AVOID
			return
		}
		fmt.Println(currentTime(), ": can't move sui to avoid capture")

		// try arata
		c.defenseARata(b, gameState.MaxLayer)
		if len(c.MoveToTarget) > 0 {
			fmt.Println(currentTime(), ": can arata koma to defense capture sui")
			c.Select = cpuselect.DEFENSE_ARATA
			return
		}
		fmt.Println(currentTime(), ": can't arata koma to defense capture sui")
		fmt.Println(currentTime(), "GG Been Capture")
		c.Select = cpuselect.BEEN_CHECKMATE
		anotherRoundOrEnd.Show = true
		return
	}

	// 嘗試俘獲
	c.tryCapture(b)
	if len(c.MoveToTarget) > 0 {
		fmt.Println(currentTime(), ": try to Capture")
		c.Select = cpuselect.TRY_CAPTURE
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
	for i, v := range c.KomaDai {
		if v.Item2 > 0 {
			probablyChoice = append(probablyChoice, []int{i})
		}
	}
	// 從可選擇列表選擇一個
	t := rand.Intn(len(probablyChoice))
	c.targetKoma = probablyChoice[t]
	if len(c.targetKoma) == 0 {
		fmt.Println(c.targetKoma)
	}
}

func currentTime() string {
	return time.Now().Format("03:04:05 PM")
}
