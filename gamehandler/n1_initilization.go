package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (g *Game) Initilization() {
	g.Font = font.ConvertToFace()
	g.GameState = gamestate.Initilization(g.Font)
}
