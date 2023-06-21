package declaresumi

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (d *DeclareSuMi) SuMiButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(constant.SUMI_BUTTON_X,
		constant.SUMI_BUTTON_Y,
		constant.SUMI_BUTTON_X+constant.SUMI_BUTTON_WIDTH,
		constant.SUMI_BUTTON_Y+constant.SUMI_BUTTON_HEIGHT))
}
