package mutiny

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

// Initilization 初始化叛變物件，設定好屬性後並回傳
//
// 參數font為繪製文字用的字型來源
func Initilization(font font.Face) Mutiny {
	var m Mutiny = Mutiny{
		Show:      false,
		TitleImg:  ebiten.NewImage(constant.MUTINY_WIDTH, constant.MUTINY_HEIGHT),
		TitleOpt:  &ebiten.DrawImageOptions{},
		MutinyImg: ebiten.NewImage(constant.MUTINY_BUTTON_WIDTH, constant.MUTINY_BUTTON_HEIGHT),
		MutinyOpt: &ebiten.DrawImageOptions{},
		CancelImg: ebiten.NewImage(constant.MUTINY_BUTTON_WIDTH, constant.MUTINY_BUTTON_HEIGHT),
		CancelOpt: &ebiten.DrawImageOptions{},
	}
	m.TitleOpt.GeoM.Scale(1.5, 1.5)
	m.TitleOpt.GeoM.Translate(constant.MUTINY_X, constant.MUTINY_Y)
	m.MutinyOpt.GeoM.Translate(constant.MUTINY_BUTTON_X, constant.MUTINY_BUTTON_Y)
	m.CancelOpt.GeoM.Translate(constant.MUTINY_BUTTON_X*1.5, constant.MUTINY_BUTTON_Y)

	text.Draw(m.TitleImg, "是否要叛變", font, 117, 32, color.Black)
	text.Draw(m.MutinyImg, "叛變", font, 30, 32, color.Black)
	text.Draw(m.CancelImg, "取消", font, 30, 32, color.Black)
	return m
}
