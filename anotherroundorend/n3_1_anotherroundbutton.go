package anotherroundorend

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// AnotherRoundButton 檢查滑鼠是否有按在"再一局"按鈕上。
//
// 若有在按鈕上則回傳true，否則false。
//
// 參數x與y為滑鼠的座標。
func (a *AnotherRoundOrEnd) AnotherRoundButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(
		constant.ANOTHER_ROUND_BUTTON_X,
		constant.ANOTHER_ROUND_BUTTON_Y,
		constant.ANOTHER_ROUND_BUTTON_X+constant.ANOTHER_ROUND_BUTTON_WIDTH,
		constant.ANOTHER_ROUND_BUTTON_Y+constant.ANOTHER_ROUND_BUTTON_HEIGHT))
}
