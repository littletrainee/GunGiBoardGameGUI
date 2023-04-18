package Board

import (
	"github.com/littletrainee/GunGiBoardGameGUI/Block"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
	"github.com/littletrainee/GunGiBoardGameGUI/Coordinate"
)

func (b *Board) Initilization() {
	var distance float32 = Const.BLOCK_SIZE + 1
	b.Blocks = make(map[Coordinate.Coordinate]Block.Block)
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			b.Blocks[Coordinate.Coordinate{X: float32(column), Y: float32(row)}] =
				Block.Block{
					Coordinate: Coordinate.Coordinate{
						X: Const.BOARD_BASE_POSITION_X + 1 + float32(column)*distance,
						Y: Const.BOARD_BASE_POSITION_Y + 1 + float32(row)*distance}}
		}
	}
}
