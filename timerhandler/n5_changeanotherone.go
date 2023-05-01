package timerhandler

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (t *TimerHandler) ChangeAnotherOne(gs *gamestate.GameState) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if t.onTimer(x, y) {
			// change anotherone's color
			t.Turn = gs.TrunToNext()
			// change to another one to count down
			// t.Turn = false
			t.Player1Timer.StopCountDown <- true
			t.Player1Timer.BackgroundColor = color.TimerPauseColor
			t.Player2Timer.CountDown()
			t.Player2Timer.BackgroundColor = color.ConfirmColor
			gs.Phase = phase.DELAY
		}
	}
}

func (t *TimerHandler) onTimer(x, y int) bool {
	rect := image.Rect(int(t.Player1Timer.CurrentCoordinate.X),
		int(t.Player1Timer.CurrentCoordinate.Y),
		int(t.Player1Timer.CurrentCoordinate.X)+int(constant.TIMER_SIZE),
		int(t.Player1Timer.CurrentCoordinate.Y)+int(constant.TIMER_SIZE))
	return image.Point{x, y}.In(rect)
}
