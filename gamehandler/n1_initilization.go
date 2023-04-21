package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/gunginotationgenerator/enum/color"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func (g *Game) Initilization() {
	g.Own = player.Player{}
	g.Opponent = player.Player{}
	g.Board = board.Board{}
	g.GameState.Level = level.BEGINNER
	g.Own.Initilization("Logan", color.BLACK, g.GameState.Level)
	g.Opponent.Initilization("Mary", color.WHITE, g.GameState.Level)
	g.Board.Initilization()
	g.Font = font.ConvertToFace()
}
