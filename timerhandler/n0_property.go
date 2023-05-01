package timerhandler

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/timer"
)

type TimerHandler struct {
	Player1Timer timer.Timer
	Player2Timer timer.Timer
	Turn         color.Gray16
}
