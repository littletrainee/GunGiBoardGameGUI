package capture

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (c *Capture) CaptureButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(constant.CAPTURE_BUTTON_POSITION_X,
		constant.CAPTURE_BUTTON_POSITION_Y,
		constant.CAPTURE_BUTTON_POSITION_X+constant.CAPTURE_BUTTON_WIDTH,
		constant.CAPTURE_BUTTON_POSITION_Y+constant.CAPTURE_BUTTON_HEIGHT))
}
