package board

import (
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
)

type Board struct {
	Blocks map[coordinate.Coordinate]block.Block
}
