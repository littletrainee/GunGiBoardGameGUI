package gamestate

import (
	"image/color"

	_level "github.com/littletrainee/GunGiBoardGameGUI/gamestate/level"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/recommendarrangement"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/selectcolor"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

type GameState struct {
	Color                      selectcolor.SelectColor
	LevelList                  []_level.Level
	Level                      level.Level
	First                      color.Gray16
	ArrangementList            []recommendarrangement.RecommendedArrangement
	RecommendedArramgement     bool
	Turn                       color.Gray16
	Phase                      phase.Phase
	ArrangementPhaseRoundCount int
	DuelingPhaseRoundCount     int
	MaxLayer                   int
}
