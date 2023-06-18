package cpu

import (
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/state"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
)

func (c *CPU) ClickClock(g *gamestate.GameState, player1Timer, player2Timer *timer.Timer, s state.State) {
	if s == state.ARRANGEMENT {
		if c.CurrentDeclareSuMiPercentage > c.DeclareSuMiTargetPercentage {
			c.Player.IsSuMi = true
			c.Player.KomaDaiBackGroundColor = color.DenyColor
			g.TurnToNext()
			player2Timer.StopCountDown <- true
			player2Timer.BackgroundColor = color.TimerPauseColor
			player1Timer.CountDown()
			player1Timer.BackgroundColor = color.ConfirmColor
		} else {
			g.TurnToNext()
			player2Timer.StopCountDown <- true
			player2Timer.BackgroundColor = color.TimerPauseColor
			player1Timer.CountDown()
			player1Timer.BackgroundColor = color.ConfirmColor
		}
	} else {
		g.TurnToNext()
		player2Timer.StopCountDown <- true
		player2Timer.BackgroundColor = color.TimerPauseColor
		player1Timer.CountDown()
		player1Timer.BackgroundColor = color.ConfirmColor
	}
}
