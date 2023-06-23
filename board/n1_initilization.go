package board

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Initilization 回傳初始化的Board物件，並設定好每個區塊在畫面上的位置、名稱與顏色。
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
					Name: image.Point{
						X: 10 - column,
						Y: row},
					Coordinate: image.Point{
						X: constant.BOARD_X + 1 + int(float32(column-1)*distance),
						Y: constant.BOARD_Y + 1 + int(float32(row-1)*distance)},
					CurrentColor: color.BoardColor}
		}
	}
	return board
}
