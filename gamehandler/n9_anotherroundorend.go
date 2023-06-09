package gamehandler

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

// AnotherRoundOrEndGame 再一局或離開遊戲畫面的判斷
func (g *Game) AnotherRoundOrEndGame() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.AnotherRoundOrEnd.AnotherRoundButton(x, y) {
			g.GameState = gamestate.Initilization(g.Font)
			g.GameState.DelayedChangePhaseTo(phase.SELECT_COLOR)
		} else if g.AnotherRoundOrEnd.EndGameButton(x, y) {
			os.Exit(0)
		}
	}
}
