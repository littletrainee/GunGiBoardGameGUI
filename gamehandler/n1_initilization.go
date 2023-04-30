package gamehandler

import (
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (g *Game) Initilization() {
	rand.Seed(time.Now().UnixNano())
	g.GameState = gamestate.Initilization()
	g.Font = font.ConvertToFace()
}
