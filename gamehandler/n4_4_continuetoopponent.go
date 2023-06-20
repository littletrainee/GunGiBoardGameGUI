package gamehandler

import "github.com/littletrainee/GunGiBoardGameGUI/color"

func (g *Game) ContinueToOpponent() {
	g.GameState.TurnToNext()
	g.Player1Timer.StopCountDown <- true
	g.Player1Timer.BackgroundColor = color.TimerPauseColor
	g.Player2Timer.CountDown()
	g.Player2Timer.BackgroundColor = color.ConfirmColor
}
