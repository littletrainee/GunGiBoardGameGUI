package timer

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// func (t *Timer) ClickTimer(gs *gamestate.GameState, pos image.Point) {
// 	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
// 		if t.onTimer(pos.X, pos.Y) {
// 			gs.TrunToNext()
// 		}
// 	}
// }

func (t *Timer) OnTimer(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(t.CurrentCoordinate.X), int(t.CurrentCoordinate.Y),
		int(t.CurrentCoordinate.X)+int(constant.TIMER_SIZE),
		int(t.CurrentCoordinate.Y)+int(constant.TIMER_SIZE)))
}
