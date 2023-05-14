package capture

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (c *Capture) Draw(screen *ebiten.Image) {
	if c.Show {
		// base
		vector.DrawFilledRect(screen, constant.CAPTURE_BASE_POSITION_X, constant.CAPTURE_BASE_POSITION_Y, constant.CAPTURE_BASE_WIDTH, constant.CAPTURE_BASE_HEIGHT, color.Black, true)
		// background
		vector.DrawFilledRect(screen, constant.CAPTURE_BASE_POSITION_X+1, constant.CAPTURE_BASE_POSITION_Y+1, constant.CAPTURE_BASE_WIDTH-2, constant.CAPTURE_BASE_HEIGHT-2, _color.BoardColor, true)
	}
	if c.CaptureBool {
		// Capture button base
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_POSITION_X, constant.CAPTURE_BUTTON_POSITION_Y, constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT, color.Black, true)
		// Capture button background
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_POSITION_X+1, constant.CAPTURE_BUTTON_POSITION_Y+1, constant.CAPTURE_BUTTON_WIDTH-2, constant.CAPTURE_BUTTON_HEIGHT-2, _color.CaptureColor, true)
		// Capture Text
		screen.DrawImage(c.CaptureImg, c.CaptureOpt)

	}
	if c.ControlBool {
		// Control Button base
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_POSITION_X*1.5, constant.CAPTURE_BUTTON_POSITION_Y, constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT, color.Black, true)
		// control button background
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_POSITION_X*1.5+1, constant.CAPTURE_BUTTON_POSITION_Y+1, constant.CAPTURE_BUTTON_WIDTH-2, constant.CAPTURE_BUTTON_HEIGHT-2, _color.DenyColor, true)

		screen.DrawImage(c.ControlImg, c.ControlOpt)
	}
}
