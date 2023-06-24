package cpu

import (
	"fmt"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/anotherroundorend"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

// DuelingPhaseSelectKoma 對弈期間的選駒，用於判斷對弈期間是否有被將軍、可否俘獲對方、移駒或新的函式
//
// 參數b為棋盤物件，gameState為當前對弈的遊戲狀態，anotherRoundOrEnd為再來一局或是離開遊戲物件
func (c *CPU) DuelingPhaseSelectKoma(b board.Board, gameState gamestate.GameState, anotherRoundOrEnd *anotherroundorend.AnotherRoundOrEnd) {
	if c.inevitableWinOpportunity(b, gameState.LevelHolder) {
		fmt.Println("GG Capture SuI")
		return
	}
	if c.inevitableLoseOpportunity(b) {

		if c.defenseCapture(b, gameState.LevelHolder) {
			fmt.Println(currentTime(), ": can capture defense")
			c.Select = cpuselect.DEFENSE_CAPTURE
			return
		}
		fmt.Println(currentTime(), ": can't capture defense")

		if c.defenseAvoid(gameState, b) {
			fmt.Println(currentTime(), ": can move sui to avoid capture")
			c.Select = cpuselect.DEFENSE_AVOID
			return
		}
		fmt.Println(currentTime(), ": can't move sui to avoid capture")

		if c.defenseARata(b, gameState) {
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
	c.RandomSelect(gameState, b)
	fmt.Println(currentTime(), ": random select koma")
	c.Select = cpuselect.RANDOM_SELECT
}

func currentTime() string {
	return time.Now().Format("03:04:05 PM")
}
