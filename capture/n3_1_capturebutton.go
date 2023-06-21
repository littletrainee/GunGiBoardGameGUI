package capture

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// CaptureButton 檢查滑鼠是否有按在"俘獲"按鈕上。
//
// 若有在按鈕上則回傳true，否則false。
//
// 參數x與y為滑鼠的座標。
func (c *Capture) CaptureButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(constant.CAPTURE_BUTTON_X,
		constant.CAPTURE_BUTTON_Y,
		constant.CAPTURE_BUTTON_X+constant.CAPTURE_BUTTON_WIDTH,
		constant.CAPTURE_BUTTON_Y+constant.CAPTURE_BUTTON_HEIGHT))
}
