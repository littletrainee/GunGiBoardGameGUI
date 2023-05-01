package timerhandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
)

func Initilization(gs gamestate.GameState) TimerHandler {
	var t TimerHandler = TimerHandler{
		Player1Timer: timer.Initilization(constant.REMAINING_TIME,
			constant.TIMER_POSITION_X+1,
			constant.TIMER_POSITION_Y+1+constant.TIMER_BASE_HEIGHT/2),
		Player2Timer: timer.Initilization(constant.REMAINING_TIME,
			constant.TIMER_POSITION_X+1,
			constant.TIMER_POSITION_Y+1),
		Turn: gs.Turn,
	}
	t.Player1Timer.BackgroundColor = color.TimerPauseColor
	t.Player2Timer.BackgroundColor = color.TimerPauseColor

	t.Player1Timer.SetGeoM(false)
	t.Player2Timer.SetGeoM(true)
	return t
}
