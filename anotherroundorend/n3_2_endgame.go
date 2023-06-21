package anotherroundorend

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// EndGameButton 檢查滑鼠位置是否有再"離開遊戲"按鈕上。
//
// 若有在按鈕上則回傳true，否則回傳false。
//
// 參數x與y為滑鼠的座標。
func (a *AnotherRoundOrEnd) EndGameButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(constant.ANOTHER_ROUND_BUTTON_POSITION_X*1.5), int(constant.ANOTHER_ROUND_BUTTON_POSITION_Y), int(constant.ANOTHER_ROUND_BUTTON_POSITION_X*1.5)+int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_POSITION_Y)+int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)))
}
