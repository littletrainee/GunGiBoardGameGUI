package capture

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (c *Capture) CaptureButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(constant.CAPTURE_BUTTON_POSITION_X),
		int(constant.CAPTURE_BUTTON_POSITION_Y),
		int(constant.CAPTURE_BUTTON_POSITION_X)+int(constant.CAPTURE_BUTTON_WIDTH),
		int(constant.CAPTURE_BUTTON_POSITION_Y)+int(constant.CAPTURE_BUTTON_HEIGHT)))
}
