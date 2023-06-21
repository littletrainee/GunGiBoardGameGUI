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
		Show:             false,
		TitleText:        "將軍",
		AnotherRoundText: "再一局",
		EndText:          "遊戲結束",
		TitleImg:         ebiten.NewImage(int(constant.ANOTHER_ROUND_OR_EXIT_WIDTH), int(constant.ANOTHER_ROUND_OR_EXIT_HEIGHT)),
		AnotherRoundImg:  ebiten.NewImage(int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)),
		EndImg:           ebiten.NewImage(int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)),
		TitleOpt:         &ebiten.DrawImageOptions{},
		AnotherRoundOpt:  &ebiten.DrawImageOptions{},
		EndOpt:           &ebiten.DrawImageOptions{},
	}

	a.TitleOpt.GeoM.Scale(1.5, 1.5)
	a.TitleOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_OR_EXIT_X), float64(constant.ANOTHER_ROUND_OR_EXIT_Y))
	a.AnotherRoundOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_BUTTON_POSITION_X), float64(constant.ANOTHER_ROUND_BUTTON_POSITION_Y))
	a.EndOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_BUTTON_POSITION_X*1.5), float64(constant.ANOTHER_ROUND_BUTTON_POSITION_Y))

	text.Draw(a.TitleImg, a.TitleText, font, 145, 32, color.Black)
	text.Draw(a.AnotherRoundImg, a.AnotherRoundText, font, 20, 32, color.Black)
	text.Draw(a.EndImg, a.EndText, font, 10, 32, color.Black)
	return a
}
