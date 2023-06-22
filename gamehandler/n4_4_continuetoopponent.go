package gamehandler

import "github.com/littletrainee/GunGiBoardGameGUI/color"

// ContinueToOpponent 對家布陣階段的繼續，若玩家點選布陣完畢按鈕，而電腦並未放置完成，則持續由電腦布陣
func (g *Game) ContinueToOpponent() {
	g.GameState.TurnToNext()
	g.Player1Timer.StopCountDown <- true
	g.Player1Timer.Background = color.TimerPauseColor
	g.Player2Timer.CountDown()
	g.Player2Timer.Background = color.ConfirmColor
}
