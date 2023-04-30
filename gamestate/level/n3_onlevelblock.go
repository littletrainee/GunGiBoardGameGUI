package level

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (l *Level) OnLevelBlock(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(l.X), int(l.Y),
		int(l.X+constant.LEVEL_BLOCK_SIZE-2),
		int(l.Y+constant.LEVEL_BLOCK_SIZE-2)))
}
