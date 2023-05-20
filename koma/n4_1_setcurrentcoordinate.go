package koma

import (
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
)

func (k *Koma) SetCurrentCoordinate(coordinate coordinate.Coordinate, shift int) {
	k.CurrentCoordinate.X = int(coordinate.X+constant.BLOCK_SIZE/2) - shift
	k.CurrentCoordinate.Y = int(coordinate.Y+constant.BLOCK_SIZE/2) - shift
}
