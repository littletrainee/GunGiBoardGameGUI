package timerhandler

func (t *TimerHandler) IsEnd() bool {
	return t.Player1Timer.RemainingTime == 0 &&
		t.Player2Timer.RemainingTime == 0
}
