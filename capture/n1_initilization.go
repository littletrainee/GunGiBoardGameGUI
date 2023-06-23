package capture

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

// Initilization 初始化Capture物件，設定好預設的顯示文字與位置並回傳。
//
// 參數font用於文字顯示得字型
func Initilization(font font.Face) Capture {
	var c Capture = Capture{
		CancelImg:                 ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		CancelOpt:                 &ebiten.DrawImageOptions{},
		CaptureImg:                ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		CaptureOpt:                &ebiten.DrawImageOptions{},
		ControlBool:               false,
		ControlOpt:                &ebiten.DrawImageOptions{},
		DisableControlImg:         ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		EnableControlImg:          ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		Show:                      false,
		TitleCaptureAndControlImg: ebiten.NewImage(476, 300),
		TitleCaptureAndControlOpt: &ebiten.DrawImageOptions{},
		TitleOnlyCaptureImg:       ebiten.NewImage(476, 300),
		TitleOnlyCaptureOpt:       &ebiten.DrawImageOptions{},
	}

	c.TitleCaptureAndControlOpt.GeoM.Scale(1.5, 1.5)
	c.TitleOnlyCaptureOpt.GeoM.Scale(1.5, 1.5)
	c.TitleCaptureAndControlOpt.GeoM.Translate(constant.CAPTURE_OR_CONTROL_X, constant.CAPTURE_OR_CONTROL_Y)
	c.TitleOnlyCaptureOpt.GeoM.Translate(constant.CAPTURE_OR_CONTROL_X, constant.CAPTURE_OR_CONTROL_Y)
	c.CaptureOpt.GeoM.Translate(constant.CAPTURE_BUTTON_X, constant.CAPTURE_BUTTON_Y)
	c.CancelOpt.GeoM.Translate(constant.CAPTURE_BUTTON_X*1.25, constant.CAPTURE_BUTTON_Y)
	c.ControlOpt.GeoM.Translate(constant.CAPTURE_BUTTON_X*1.5, constant.CAPTURE_BUTTON_Y)

	text.Draw(c.TitleCaptureAndControlImg, "要俘獲還是控制", font, 90, 32, color.Black)
	text.Draw(c.TitleOnlyCaptureImg, "是否要俘獲", font, 110, 32, color.Black)
	text.Draw(c.CaptureImg, "俘獲", font, 30, 32, color.Black)
	text.Draw(c.CancelImg, "取消", font, 30, 32, color.Black)
	text.Draw(c.EnableControlImg, "控制", font, 30, 32, color.Black)
	text.Draw(c.DisableControlImg, "控制", font, 30, 32, _color.DisableButtonColor)
	return c
}
