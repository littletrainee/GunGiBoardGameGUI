package gamestate

import (
	"github.com/littletrainee/gunginotationgenerator/enum/color"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

type GameState struct {
	Level level.Level
	Turn  color.Color
	Round int
}
