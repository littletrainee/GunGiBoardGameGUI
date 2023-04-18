package GameHolder

import (
	"github.com/littletrainee/GunGiBoardGameGUI/Board"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
	"github.com/littletrainee/GunGiBoardGameGUI/Font"
	"github.com/littletrainee/GunGiBoardGameGUI/Player"
)

func (g *Game) Initilization() {
	g.Own = Player.Player{}
	g.Opponent = Player.Player{}
	g.Board = Board.Board{}
	g.GameState.Level = Const.Advanced
	g.Own.Initilization("Logan", Const.Black, g.GameState.Level)
	g.Opponent.Initilization("Mary", Const.White, g.GameState.Level)
	g.Board.Initilization()
	g.Font = Font.ConvertToFace()
}
