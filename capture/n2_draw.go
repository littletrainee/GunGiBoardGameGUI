package capture

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製Capture彈出視窗
//
// 參數screen為要繪製的目標視窗
func (c *Capture) Draw(screen *ebiten.Image) {
	if c.Show {
		// base
		vector.DrawFilledRect(screen, constant.CAPTURE_OR_CONTROL_X, constant.CAPTURE_OR_CONTROL_Y, constant.CAPTURE_OR_CONTROL_WIDTH, constant.CAPTURE_OR_CONTROL_HEIGHT, color.Black, true)
		// background
		vector.DrawFilledRect(screen, constant.CAPTURE_OR_CONTROL_X+1, constant.CAPTURE_OR_CONTROL_Y+1, constant.CAPTURE_OR_CONTROL_WIDTH-2, constant.CAPTURE_OR_CONTROL_HEIGHT-2, _color.BoardColor, true)

		// Capture button
		// base
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X, constant.CAPTURE_BUTTON_Y, constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT, color.Black, true)
		// background
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X+1, constant.CAPTURE_BUTTON_Y+1, constant.CAPTURE_BUTTON_WIDTH-2, constant.CAPTURE_BUTTON_HEIGHT-2, _color.CaptureColor, true)
		// Text
		screen.DrawImage(c.CaptureImg, c.CaptureOpt)

		// Cancel button
		// base
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X*1.25, constant.CAPTURE_BUTTON_Y, constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT, color.Black, true)
		// background
		vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X*1.25+1, constant.CAPTURE_BUTTON_Y+1, constant.CAPTURE_BUTTON_WIDTH-2, constant.CAPTURE_BUTTON_HEIGHT-2, _color.DenyColor, true)
		// Text
		screen.DrawImage(c.CancelImg, c.CancelOpt)

		// Control button
		if c.ControlBool {
			screen.DrawImage(c.TitleCaptureAndControlImg, c.TitleCaptureAndControlOpt)
			// base
			vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X*1.5, constant.CAPTURE_BUTTON_Y, constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT, color.Black, true)
			// background
			vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X*1.5+1, constant.CAPTURE_BUTTON_Y+1, constant.CAPTURE_BUTTON_WIDTH-2, constant.CAPTURE_BUTTON_HEIGHT-2, _color.ConfirmColor, true)
			// Text
			screen.DrawImage(c.EnableControlImg, c.ControlOpt)
		} else {
			screen.DrawImage(c.TitleOnlyCaptureImg, c.TitleOnlyCaptureOpt)
			// base
			vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X*1.5, constant.CAPTURE_BUTTON_Y, constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT, color.Black, true)
			// background
			vector.DrawFilledRect(screen, constant.CAPTURE_BUTTON_X*1.5+1, constant.CAPTURE_BUTTON_Y+1, constant.CAPTURE_BUTTON_WIDTH-2, constant.CAPTURE_BUTTON_HEIGHT-2, _color.BoardColor, true)
			// Text
			screen.DrawImage(c.DisableControlImg, c.ControlOpt)

		}
	}
}
