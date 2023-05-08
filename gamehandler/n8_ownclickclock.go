package gamehandler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

func (g *Game) OwnClickClock() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.Player1Timer.OnTimer(x, y) {
			// change anotherone's color
			g.GameState.TurnToNext()
			g.Player1Timer.StopCountDown <- true
			g.Player1Timer.BackgroundColor = color.TimerPauseColor
			g.Player2Timer.CountDown()
			g.Player2Timer.BackgroundColor = color.ConfirmColor
			g.delayedChangePhaseTo(phase.DUELING_PHASE_SELECT_KOMA)
		}
	}
}
