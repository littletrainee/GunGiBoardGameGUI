package board

import (
	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
)

func (b *Board) Initilization() {
	var distance float32 = constant.BLOCK_SIZE + 1
	b.Blocks = make(map[coordinate.Coordinate]block.Block)
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			b.Blocks[coordinate.Coordinate{X: float32(column), Y: float32(row)}] =
				block.Block{
					Coordinate: coordinate.Coordinate{
						X: constant.BOARD_BASE_POSITION_X + 1 + float32(column)*distance,
						Y: constant.BOARD_BASE_POSITION_Y + 1 + float32(row)*distance}}
		}
	}
}
