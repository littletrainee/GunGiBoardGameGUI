package timerhandler

func (t *TimerHandler) Count() {
	t.Player1Timer.CountDown()
	t.Player2Timer.CountDown()
}
