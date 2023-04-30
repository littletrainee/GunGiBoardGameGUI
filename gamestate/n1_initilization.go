package gamestate

import (
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/selectcolor"
)

func Initilization() GameState {
	return GameState{Phase: phase.SELECT_COLOR, Color: selectcolor.Initilization()}
	// g.Level = gameLevel
	// g.First = firstColor
	// g.Turn = firstColor
	// if gameLevel == level.ADVANCED {
	// 	g.MaxLayer = 3
	// } else {
	// 	g.MaxLayer = 2
	// }
	// g.Phase = phase.LEVEL_SELECT
	// g.ArrangementPhaseRoundCount = 1
	// g.DuelingPhaseRoundCount = 0
}
