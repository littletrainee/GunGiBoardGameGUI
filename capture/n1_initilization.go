package capture

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func Initilization(font font.Face) Capture {

	var c Capture = Capture{
		Show:        false,
		CaptureBool: false,
		ControlBool: false,
		CaptureText: "俘獲",
		ControlText: "控制",
		CaptureImg:  ebiten.NewImage(int(constant.CAPTURE_BUTTON_WIDTH), int(constant.CAPTURE_BUTTON_HEIGHT)),
		ControlImg:  ebiten.NewImage(int(constant.CAPTURE_BUTTON_WIDTH), int(constant.CAPTURE_BUTTON_HEIGHT)),
		CaptureOpt:  &ebiten.DrawImageOptions{},
		ControlOpt:  &ebiten.DrawImageOptions{},
	}
	c.CaptureOpt.GeoM.Translate(float64(constant.CAPTURE_BUTTON_POSITION_X), float64(constant.CAPTURE_BUTTON_POSITION_Y))
	c.ControlOpt.GeoM.Translate(float64(constant.CAPTURE_BUTTON_POSITION_X*1.5), float64(constant.CAPTURE_BUTTON_POSITION_Y))
	text.Draw(c.CaptureImg, c.CaptureText, font, 30, 32, color.Black)
	text.Draw(c.ControlImg, c.ControlText, font, 30, 32, color.Black)
	return c
}
