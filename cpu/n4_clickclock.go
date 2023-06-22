package cpu

import (
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/state"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
)

// ClickClock 按下棋鐘，電腦移居後的按下棋鐘，將控制權轉為玩家1
//
// 參數g為遊戲當前的狀態，player1Timer與player2Timer為兩個玩家各自的棋鐘，s為遊戲期間式布陣還是對弈
func (c *CPU) ClickClock(g *gamestate.GameState, player1Timer, player2Timer *timer.Timer, s state.State) {
	if s == state.ARRANGEMENT {
		if c.CurrentDeclareSuMiPercentage > c.DeclareSuMiTargetPercentage {
			c.Player.IsSuMi = true
			c.Player.KomaDaiBackground = color.DenyColor
			g.TurnToNext()
			player2Timer.StopCountDown <- true
			player2Timer.Background = color.TimerPauseColor
			player1Timer.CountDown()
			player1Timer.Background = color.ConfirmColor
		} else {
			g.TurnToNext()
			player2Timer.StopCountDown <- true
			player2Timer.Background = color.TimerPauseColor
			player1Timer.CountDown()
			player1Timer.Background = color.ConfirmColor
		}
	} else {
		g.TurnToNext()
		player2Timer.StopCountDown <- true
		player2Timer.Background = color.TimerPauseColor
		player1Timer.CountDown()
		player1Timer.Background = color.ConfirmColor
	}
}
