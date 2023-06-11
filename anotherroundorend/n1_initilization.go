package anotherroundorend

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func Initilization(font font.Face) AnotherRoundOrEnd {
	var a AnotherRoundOrEnd = AnotherRoundOrEnd{
		Show:             false,
		AnotherRoundText: "再一局",
		EndText:          "遊戲結束",
		AnotherRoundImg:  ebiten.NewImage(int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)),
		EndImg:           ebiten.NewImage(int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)),
		AnotherRoundOpt:  &ebiten.DrawImageOptions{},
		EndOpt:           &ebiten.DrawImageOptions{},
	}
	a.AnotherRoundOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_BUTTON_POSITION_X), float64(constant.ANOTHER_ROUND_BUTTON_POSITION_Y))
	a.EndOpt.GeoM.Translate(float64(constant.ANOTHER_ROUND_BUTTON_POSITION_X*1.5), float64(constant.ANOTHER_ROUND_BUTTON_POSITION_Y))
	text.Draw(a.AnotherRoundImg, a.AnotherRoundText, font, 20, 32, color.Black)
	text.Draw(a.EndImg, a.EndText, font, 10, 32, color.Black)
	return a
}
