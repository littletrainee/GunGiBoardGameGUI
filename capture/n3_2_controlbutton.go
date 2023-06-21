package capture

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// ControlButton 檢查滑鼠是否有按在"控制"按鈕上。
//
// 若有在按鈕上則回傳true，否則false。
//
// 參數x與y為滑鼠的座標。
func (c *Capture) ControlButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(constant.CAPTURE_BUTTON_X*1.5),
		int(constant.CAPTURE_BUTTON_Y),
		int(constant.CAPTURE_BUTTON_X*1.5)+int(constant.CAPTURE_BUTTON_WIDTH),
		int(constant.CAPTURE_BUTTON_Y)+int(constant.CAPTURE_BUTTON_HEIGHT)))
}
