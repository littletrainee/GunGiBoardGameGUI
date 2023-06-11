package anotherroundorend

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (a *AnotherRoundOrEnd) AnotherRoundButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(constant.ANOTHER_ROUND_BUTTON_POSITION_X), int(constant.ANOTHER_ROUND_BUTTON_POSITION_Y), int(constant.ANOTHER_ROUND_BUTTON_POSITION_X)+int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_POSITION_Y)+int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)))
}
