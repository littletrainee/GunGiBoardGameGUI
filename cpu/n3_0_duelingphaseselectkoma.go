package cpu

import (
	"log"

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
		log.Printf("自家被將死\n\n")
		return
	}
	if c.inevitableLoseOpportunity(b, gameState.LevelHolder) {

		if c.defenseCapture(b, gameState.LevelHolder) {
			log.Printf("可以用俘獲阻止被將軍\n%s:%v\n\n", "c.MoveToTarget", c.MoveToTarget)
			c.Select = cpuselect.DEFENSE_CAPTURE
			return
		}
		log.Printf("無法用俘獲阻止被將軍\n\n")

		if c.defenseAvoid(gameState, b) {
			log.Printf("可以用移動帥避免被將軍\n%s:%v\n\n", "c.MoveToTarget", c.MoveToTarget)
			c.Select = cpuselect.DEFENSE_AVOID
			return
		}
		log.Printf("無法移動帥避免被將軍\n\n")
		log.Printf("對家被將死\n\n")
		c.Select = cpuselect.BEEN_CHECKMATE
		anotherRoundOrEnd.Show = true
		return
	}

	if c.tryCapture(b) {
		log.Printf("嘗試俘獲\n%s:%v\n\n", "c.MoveToTarget", c.MoveToTarget)
		c.Select = cpuselect.TRY_CAPTURE
		return
	}
	c.RandomSelect(gameState, b)
	log.Printf("隨機選擇駒與移動位置\n%s:%v\n\n", "c.MoveToTarget", c.MoveToTarget)
	c.Select = cpuselect.RANDOM_SELECT
}
