package board

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
)

type Board struct {
	Blocks map[image.Point]block.Block
}
