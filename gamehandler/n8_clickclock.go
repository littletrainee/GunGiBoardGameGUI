package gamehandler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/state"
)

// ClickClock 棋鐘按鈕按下的判斷
func (g *Game) ClickClock() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// 若按下地棋鐘的位置則換對家
		if g.Player1Timer.OnTimer(x, y) {
			// change anotherone's color
			g.GameState.TurnToNext()
			g.Player1Timer.StopCountDown <- true
			g.Player1Timer.Background = color.TimerPauseColor
			g.Player2Timer.CountDown()
			g.Player2Timer.Background = color.ConfirmColor
			g.GameState.DelayedChangePhaseTo(phase.SELECT_KOMA)
		} else if g.DeclareSuMi.SuMiButton(x, y) && g.CurrentState == state.ARRANGEMENT {
			g.Player1.IsSuMi = true
			g.DeclareSuMi.Show = false
			g.Player1.KomaDaiBackground = color.DenyColor
			g.GameState.TurnToNext()
			g.Player1Timer.StopCountDown <- true
			g.Player1Timer.Background = color.TimerPauseColor
			g.Player2Timer.CountDown()
			g.Player2Timer.Background = color.ConfirmColor
			g.GameState.DelayedChangePhaseTo(phase.SELECT_KOMA)
		}
	}
}
