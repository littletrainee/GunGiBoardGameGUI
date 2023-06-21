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
		Show:              false,
		ControlBool:       false,
		CaptureText:       "俘獲",
		CancelText:        "取消",
		ControlText:       "控制",
		CaptureImg:        ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		CancelImg:         ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		EnableControlImg:  ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		DisableControlImg: ebiten.NewImage(constant.CAPTURE_BUTTON_WIDTH, constant.CAPTURE_BUTTON_HEIGHT),
		CaptureOpt:        &ebiten.DrawImageOptions{},
		CancelOpt:         &ebiten.DrawImageOptions{},
		ControlOpt:        &ebiten.DrawImageOptions{},
	}
	c.CaptureOpt.GeoM.Translate(constant.CAPTURE_BUTTON_X, constant.CAPTURE_BUTTON_Y)
	c.CancelOpt.GeoM.Translate(constant.CAPTURE_BUTTON_X*1.25, constant.CAPTURE_BUTTON_Y)
	c.ControlOpt.GeoM.Translate(constant.CAPTURE_BUTTON_X*1.5, constant.CAPTURE_BUTTON_Y)
	text.Draw(c.CaptureImg, c.CaptureText, font, 30, 32, color.Black)
	text.Draw(c.CancelImg, c.CancelText, font, 30, 32, color.Black)
	text.Draw(c.EnableControlImg, c.ControlText, font, 30, 32, color.Black)
	text.Draw(c.DisableControlImg, c.CancelText, font, 30, 32, _color.DisableButtonColor)
	return c
}
