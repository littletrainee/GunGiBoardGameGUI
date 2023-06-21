package capture

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// CancelButton 檢查滑鼠是否有按在"取消"按鈕上。
//
// 若有在按鈕上則回傳true，否則false。
//
// 參數x與y為滑鼠的座標。
func (c *Capture) CancelButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(constant.CAPTURE_BUTTON_X*1.25),
		int(constant.CAPTURE_BUTTON_Y),
		int(constant.CAPTURE_BUTTON_X*1.25)+int(constant.CAPTURE_BUTTON_WIDTH),
		int(constant.CAPTURE_BUTTON_Y)+int(constant.CAPTURE_BUTTON_HEIGHT)))
}
