package gamestate

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
	_level "github.com/littletrainee/GunGiBoardGameGUI/gamestate/level"
	ron "github.com/littletrainee/GunGiBoardGameGUI/gamestate/recommendarrangement"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/selectcolor"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

// 遊戲總體狀態
type GameState struct {
	Color                      selectcolor.SelectColor             // 供玩家選擇顏色的選項
	LevelList                  []_level.Level                      // 供玩家選擇階級的選項
	Level                      level.Level                         // 玩家選擇的階級
	First                      color.Gray16                        // 先手的顏色
	RecommendOrManual          []ron.RecommendedArrangement        // 布陣的選擇
	ArramgementBy              arrangementselect.ArrangementSelect // 在入門或出擊的推薦或是自訂布陣的選擇
	Turn                       color.Gray16                        // 換誰
	Phase                      phase.Phase                         // 程式當前的階段
	ArrangementPhaseRoundCount int                                 // 布陣階段的巡數記數
	DuelingPhaseRoundCount     int                                 // 對弈階段的巡數記數
	MaxLayer                   int                                 // 最大段位
}
