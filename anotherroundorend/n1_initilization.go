package anotherroundorend

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/checkmate"
	"golang.org/x/image/font"
)

// Initilization 初始化AnotherRoundOrEnd物件，設定好預設的顯示文字與位置並回傳。
//
// 參數font用於文字顯示得字型
func Initilization(font font.Face) AnotherRoundOrEnd {
	var a AnotherRoundOrEnd = AnotherRoundOrEnd{
		Show:                 false,
		CheckMateEnumeration: checkmate.NONE,
		TitleCheckMateImg:    ebiten.NewImage(constant.ANOTHER_ROUND_OR_EXIT_WIDTH, constant.ANOTHER_ROUND_OR_EXIT_HEIGHT),
		TitleCheckMateOpt:    &ebiten.DrawImageOptions{},
		TitleBeCheckMateImg:  ebiten.NewImage(constant.ANOTHER_ROUND_OR_EXIT_WIDTH, constant.ANOTHER_ROUND_OR_EXIT_HEIGHT),
		TitleBeCheckMateOpt:  &ebiten.DrawImageOptions{},
		AnotherRoundImg:      ebiten.NewImage(constant.ANOTHER_ROUND_BUTTON_WIDTH, constant.ANOTHER_ROUND_BUTTON_HEIGHT),
		AnotherRoundOpt:      &ebiten.DrawImageOptions{},
		EndImg:               ebiten.NewImage(constant.ANOTHER_ROUND_BUTTON_WIDTH, constant.ANOTHER_ROUND_BUTTON_HEIGHT),
		EndOpt:               &ebiten.DrawImageOptions{},
	}

	a.TitleCheckMateOpt.GeoM.Scale(1.5, 1.5)
	a.TitleCheckMateOpt.GeoM.Translate(constant.ANOTHER_ROUND_OR_EXIT_X, constant.ANOTHER_ROUND_OR_EXIT_Y)
	a.TitleBeCheckMateOpt.GeoM.Scale(1.5, 1.5)
	a.TitleBeCheckMateOpt.GeoM.Translate(constant.ANOTHER_ROUND_OR_EXIT_X, constant.ANOTHER_ROUND_OR_EXIT_Y)
	a.AnotherRoundOpt.GeoM.Translate(constant.ANOTHER_ROUND_BUTTON_X, constant.ANOTHER_ROUND_BUTTON_Y)
	a.EndOpt.GeoM.Translate(constant.ANOTHER_ROUND_BUTTON_X*1.5, constant.ANOTHER_ROUND_BUTTON_Y)

	text.Draw(a.TitleCheckMateImg, "將軍", font, 145, 32, color.Black)
	text.Draw(a.TitleBeCheckMateImg, "被將軍", font, 132, 32, color.Black)
	text.Draw(a.AnotherRoundImg, "再一局", font, 20, 32, color.Black)
	text.Draw(a.EndImg, "遊戲結束", font, 10, 32, color.Black)
	return a
}
