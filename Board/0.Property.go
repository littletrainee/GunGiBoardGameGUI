package Board

import (
	"github.com/littletrainee/GunGiBoardGameGUI/Block"
	"github.com/littletrainee/GunGiBoardGameGUI/Coordinate"
)

type Board struct {
	Blocks map[Coordinate.Coordinate]Block.Block
}
