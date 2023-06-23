package anotherroundorend

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製AnotherRoundOrEnd彈出視窗
//
// 參數screen為要繪製的目標視窗
func (a *AnotherRoundOrEnd) Draw(screen *ebiten.Image) {
	if a.Show {
		// base
		vector.DrawFilledRect(screen,
			constant.ANOTHER_ROUND_OR_EXIT_X, constant.ANOTHER_ROUND_OR_EXIT_Y,
			constant.ANOTHER_ROUND_OR_EXIT_WIDTH, constant.ANOTHER_ROUND_OR_EXIT_HEIGHT, color.Black, true)
		// background
		vector.DrawFilledRect(screen,
			constant.ANOTHER_ROUND_OR_EXIT_X+1, constant.ANOTHER_ROUND_OR_EXIT_Y+1,
			constant.ANOTHER_ROUND_OR_EXIT_WIDTH-2, constant.ANOTHER_ROUND_OR_EXIT_HEIGHT-2, _color.BoardColor, true)
		screen.DrawImage(a.TitleImg, a.TitleOpt)

		// another round button base
		vector.DrawFilledRect(screen,
			float32(constant.ANOTHER_ROUND_BUTTON_X), float32(constant.ANOTHER_ROUND_BUTTON_Y), constant.ANOTHER_ROUND_BUTTON_WIDTH, constant.ANOTHER_ROUND_BUTTON_HEIGHT, color.Black, true)
		// another round button background
		vector.DrawFilledRect(screen,
			constant.ANOTHER_ROUND_BUTTON_X+1, constant.ANOTHER_ROUND_BUTTON_Y+1,
			constant.ANOTHER_ROUND_BUTTON_WIDTH-2, constant.ANOTHER_ROUND_BUTTON_HEIGHT-2, _color.CaptureColor, true)
		screen.DrawImage(a.AnotherRoundImg, a.AnotherRoundOpt)

		// end game button base
		vector.DrawFilledRect(screen,
			constant.ANOTHER_ROUND_BUTTON_X*1.5, constant.ANOTHER_ROUND_BUTTON_Y,
			constant.ANOTHER_ROUND_BUTTON_WIDTH, constant.ANOTHER_ROUND_BUTTON_HEIGHT, color.Black, true)
		// end game button background
		vector.DrawFilledRect(screen,
			constant.ANOTHER_ROUND_BUTTON_X*1.5+1, constant.ANOTHER_ROUND_BUTTON_Y+1,
			constant.ANOTHER_ROUND_BUTTON_WIDTH-2, constant.ANOTHER_ROUND_BUTTON_HEIGHT-2, _color.DenyColor, true)
		screen.DrawImage(a.EndImg, a.EndOpt)
	}
}
