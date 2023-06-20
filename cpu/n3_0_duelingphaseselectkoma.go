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

		if c.defenseCapture(b) {
			fmt.Println(currentTime(), ": can capture defense")
			c.Select = cpuselect.DEFENSE_CAPTURE
			return
		}
		fmt.Println(currentTime(), ": can't capture defense")

		if c.defenseAvoid(b) {
			fmt.Println(currentTime(), ": can move sui to avoid capture")
			c.Select = cpuselect.DEFENSE_AVOID
			return
		}
		fmt.Println(currentTime(), ": can't move sui to avoid capture")

		if c.defenseARata(b, gameState.MaxLayer) {
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

	if c.tryCapture(b) {
		fmt.Println(currentTime(), ": try to Capture")
		c.Select = cpuselect.TRY_CAPTURE
		return
	}

	var probablyChoice [][]int
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
}

func currentTime() string {
	return time.Now().Format("03:04:05 PM")
}
