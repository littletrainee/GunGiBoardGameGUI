package declaresumi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (d *DeclareSuMi) Draw(screen *ebiten.Image) {
	if d.Show {
		vector.DrawFilledRect(screen, constant.SUMI_BUTTON_POSITION_X, constant.SUMI_BUTTON_POSITION_Y, constant.SUMI_BUTTON_WIDTH, constant.SUMI_BUTTON_HEIGHT, color.Black, true)
		vector.DrawFilledRect(screen, constant.SUMI_BUTTON_POSITION_X+1, constant.SUMI_BUTTON_POSITION_Y+1, constant.SUMI_BUTTON_WIDTH-2, constant.SUMI_BUTTON_HEIGHT-2, _color.CaptureColor, true)

		screen.DrawImage(d.SuMiImg, d.SuMiOpt)
	}
}
