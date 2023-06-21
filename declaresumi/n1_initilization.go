package declaresumi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func Initilization(font font.Face) DeclareSuMi {
	var d DeclareSuMi = DeclareSuMi{
		Show:     true,
		SuMiText: "布鎮完畢",
		SuMiImg:  ebiten.NewImage(int(constant.SUMI_BUTTON_WIDTH), int(constant.SUMI_BUTTON_HEIGHT)),
		SuMiOpt:  &ebiten.DrawImageOptions{},
	}
	d.SuMiOpt.GeoM.Translate(constant.SUMI_BUTTON_X, constant.SUMI_BUTTON_Y)
	text.Draw(d.SuMiImg, d.SuMiText, font, 10, 55, color.Black)
	return d
}
