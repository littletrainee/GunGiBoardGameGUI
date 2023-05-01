package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (g *Game) Initilization() {
	g.GameState = gamestate.Initilization()
	g.Font = font.ConvertToFace()
}
