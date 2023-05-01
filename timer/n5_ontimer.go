package timer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (t *Timer) ClickTimer(gs *gamestate.GameState, pos image.Point) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if t.onTimer(pos.X, pos.Y) {
			gs.TrunToNext()
		}
	}
}

func (t *Timer) onTimer(x, y int) bool {
	rect := image.Rect(int(t.CurrentCoordinate.X), int(t.CurrentCoordinate.Y),
		int(t.CurrentCoordinate.X)+int(constant.TIMER_SIZE),
		int(t.CurrentCoordinate.Y)+int(constant.TIMER_SIZE))
	return image.Point{x, y}.In(rect)
}
