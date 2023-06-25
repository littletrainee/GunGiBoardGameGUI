package mutiny

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// CancelButton 檢查滑鼠位置是否在取消按鈕上，若在按鈕上回傳true，否則回傳false
//
// 參數x與y為滑鼠的座標
func (m *Mutiny) CancelButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(
		constant.MUTINY_BUTTON_X*1.5,
		constant.MUTINY_BUTTON_Y,
		constant.MUTINY_BUTTON_X*1.5+constant.MUTINY_BUTTON_WIDTH,
		constant.MUTINY_BUTTON_Y+constant.MUTINY_BUTTON_HEIGHT))
}
