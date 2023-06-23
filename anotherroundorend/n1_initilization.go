package anotherroundorend

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

// Initilization 初始化AnotherRoundOrEnd物件，設定好預設的顯示文字與位置並回傳。
//
// 參數font用於文字顯示得字型
func Initilization(font font.Face) AnotherRoundOrEnd {
	var a AnotherRoundOrEnd = AnotherRoundOrEnd{
		AnotherRoundImg: ebiten.NewImage(int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)),
		AnotherRoundOpt: &ebiten.DrawImageOptions{},
		EndImg:          ebiten.NewImage(int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)),
		EndOpt:          &ebiten.DrawImageOptions{},
		Show:            false,
		TitleImg:        ebiten.NewImage(int(constant.ANOTHER_ROUND_OR_EXIT_WIDTH), int(constant.ANOTHER_ROUND_OR_EXIT_HEIGHT)),
		TitleOpt:        &ebiten.DrawImageOptions{},
	}

	a.TitleOpt.GeoM.Scale(1.5, 1.5)
	a.TitleOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_OR_EXIT_X), float64(constant.ANOTHER_ROUND_OR_EXIT_Y))
	a.AnotherRoundOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_BUTTON_X), float64(constant.ANOTHER_ROUND_BUTTON_Y))
	a.EndOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_BUTTON_X*1.5), float64(constant.ANOTHER_ROUND_BUTTON_Y))

	text.Draw(a.TitleImg, "將軍", font, 145, 32, color.Black)
	text.Draw(a.AnotherRoundImg, "再一局", font, 20, 32, color.Black)
	text.Draw(a.EndImg, "遊戲結束", font, 10, 32, color.Black)
	return a
}
