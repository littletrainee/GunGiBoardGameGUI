package mutiny

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// MutinyButton 叛變按鈕的判斷，判斷滑鼠是否在按鈕上，若在按鈕上回傳true，否則回傳false
//
// 參數x與y為滑鼠的座標
func (m *Mutiny) MutinyButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(
		constant.MUTINY_BUTTON_X,
		constant.MUTINY_BUTTON_Y,
		constant.MUTINY_BUTTON_X+constant.MUTINY_BUTTON_WIDTH,
		constant.MUTINY_BUTTON_Y+constant.MUTINY_HEIGHT))
}
