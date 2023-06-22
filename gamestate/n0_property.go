// 當前遊戲總體狀態的物件
package gamestate

import (
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/arrangementholder"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/colorholder"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

// 當前遊戲總體狀態
type GameState struct {
	ColorHolder                colorholder.ColorHolder             // 供玩家選擇顏色的選項
	LevelHolder                levelholder.LevelHolder             // 供玩家選擇階級的選項
	ArrangementHolder          arrangementholder.ArrangementHolder // 當玩家選擇入門或是初級階級時，供玩家選擇推薦或自訂布陣的選項
	Phase                      phase.Phase                         // 程式當前的階段
	ArrangementPhaseRoundCount int                                 // 布陣階段的巡數記數
	DuelingPhaseRoundCount     int                                 // 對弈階段的巡數記數
}
