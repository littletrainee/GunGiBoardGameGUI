package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/font"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

// Initilization 初始化Game物件，預先處理字型與初始化整體的遊戲狀態。
func (g *Game) Initilization() {
	g.Font = font.ConvertToFace()
	g.GameState = gamestate.Initilization(g.Font)
}
