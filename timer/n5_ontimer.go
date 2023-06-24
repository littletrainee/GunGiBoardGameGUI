package timer

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// OnTimer 判斷滑鼠是否在棋鐘圖像上，若在上面則回傳true，否則回傳false
//
// 參數x與y為滑鼠的位置
func (t *Timer) OnTimer(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(t.CurrentCoordinate.X), int(t.CurrentCoordinate.Y),
		int(t.CurrentCoordinate.X)+int(constant.TIMER_SIZE),
		int(t.CurrentCoordinate.Y)+int(constant.TIMER_SIZE)))
}
