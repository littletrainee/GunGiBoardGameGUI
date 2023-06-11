package gamehandler

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (g *Game) AnotherRoundOrEndGame() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.AnotherRoundOrEnd.AnotherRoundButton(x, y) {
			g.GameState = gamestate.Initilization(g.Font)
			g.delayedChangePhaseTo(phase.SELECT_COLOR)
		} else if g.AnotherRoundOrEnd.EndGame(x, y) {
			os.Exit(0)
		}
	}
}
