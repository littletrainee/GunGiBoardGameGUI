package block

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (b *Block) OnBlock(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(int(b.Coordinate.X),
		int(b.Coordinate.Y), int(b.Coordinate.X)+int(constant.BLOCK_SIZE),
		int(b.Coordinate.Y)+int(constant.BLOCK_SIZE)))
}
