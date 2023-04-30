package timerhandler

import "github.com/littletrainee/GunGiBoardGameGUI/timer"

type TimerHandler struct {
	Player1Timer timer.Timer
	Player2Timer timer.Timer
	// true is player1 false is player2
	Turn bool
}
