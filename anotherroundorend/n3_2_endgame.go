package anotherroundorend

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (a *AnotherRoundOrEnd) EndGame(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(constant.ANOTHER_ROUND_BUTTON_POSITION_X*1.5), int(constant.ANOTHER_ROUND_BUTTON_POSITION_Y), int(constant.ANOTHER_ROUND_BUTTON_POSITION_X*1.5)+int(constant.ANOTHER_ROUND_BUTTON_WIDTH), int(constant.ANOTHER_ROUND_BUTTON_POSITION_Y)+int(constant.ANOTHER_ROUND_BUTTON_HEIGHT)))
}
