package declaresumi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

// Initilization 初始化DeclareSuMi物件，設定好預設的顯示文字與位置並回傳。
//
// 參數font用於文字顯示得字型
func Initilization(font font.Face) DeclareSuMi {
	var d DeclareSuMi = DeclareSuMi{
		Show:    true,
		SuMiImg: ebiten.NewImage(int(constant.SUMI_BUTTON_WIDTH), int(constant.SUMI_BUTTON_HEIGHT)),
		SuMiOpt: &ebiten.DrawImageOptions{},
	}
	d.SuMiOpt.GeoM.Translate(constant.SUMI_BUTTON_X, constant.SUMI_BUTTON_Y)
	text.Draw(d.SuMiImg, "布鎮完畢", font, 10, 55, color.Black)
	return d
}
