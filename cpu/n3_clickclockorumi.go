package cpu

import (
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/timerhandler"
)

func (c *CPU) ClickClockOrSuMi(t *timerhandler.TimerHandler, gs *gamestate.GameState) {
	if c.DeclareSuMiPercentage > c.target {
		c.Player.IsSuMi = true
	} else {
		gs.TrunToNext()
		t.Turn = gs.Turn
		t.Player2Timer.StopCountDown <- true
		t.Player2Timer.BackgroundColor = color.TimerPauseColor
		t.Player1Timer.CountDown()
		t.Player1Timer.BackgroundColor = color.ConfirmColor
		c.DeclareSuMiPercentage += c.PercentagePhase
	}
}
