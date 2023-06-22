package declaresumi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製DeclareSumi彈出視窗
//
// 參數screen為要繪製的目標視窗
func (d *DeclareSuMi) Draw(screen *ebiten.Image) {
	if d.Show {
		vector.DrawFilledRect(screen, constant.SUMI_BUTTON_X, constant.SUMI_BUTTON_Y, constant.SUMI_BUTTON_WIDTH, constant.SUMI_BUTTON_HEIGHT, color.Black, true)
		vector.DrawFilledRect(screen, constant.SUMI_BUTTON_X+1, constant.SUMI_BUTTON_Y+1, constant.SUMI_BUTTON_WIDTH-2, constant.SUMI_BUTTON_HEIGHT-2, _color.CaptureColor, true)

		screen.DrawImage(d.SuMiImg, d.SuMiOpt)
	}
}
