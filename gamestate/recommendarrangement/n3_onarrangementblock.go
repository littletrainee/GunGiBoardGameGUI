package recommendarrangement

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (i *RecommendedArrangement) OnAskArrangement(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(i.X), int(i.Y),
		int(i.X+constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE-2),
		int(i.Y+constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE-2)))
}
