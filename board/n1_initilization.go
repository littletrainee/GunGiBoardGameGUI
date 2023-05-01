package board

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
	"github.com/littletrainee/pair"
)

func Initilization() Board {
	var (
		distance float32 = constant.BLOCK_SIZE + 1
		board    Board   = Board{}
	)

	board.Blocks = make(map[image.Point]block.Block)
	for row := 1; row < 10; row++ {
		for column := 1; column < 10; column++ {
			board.Blocks[image.Point{X: 10 - column, Y: row}] =
				block.Block{
					Name: pair.Pair[int, int]{
						Item1: 10 - column,
						Item2: row},
					Coordinate: coordinate.Coordinate{
						X: constant.BOARD_BASE_POSITION_X + 1 + float32(column-1)*distance,
						Y: constant.BOARD_BASE_POSITION_Y + 1 + float32(row-1)*distance},
					CurrentColor: color.BoardColor}
		}
	}
	return board
}
