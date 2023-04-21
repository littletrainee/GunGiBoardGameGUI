package boardandkomatai

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func Board(screen *ebiten.Image) {
	// Base
	vector.DrawFilledRect(screen, constant.BOARD_BASE_POSITION_X,
		constant.BOARD_BASE_POSITION_Y, constant.BOARD_BASE_SIZE_WIDTH, constant.BOARD_BASE_SIZE_WIDTH,
		color.Black, true)
	// Row 1
	for i := 0; i < 9; i++ {
		row(screen, float32(i))
	}
}

func row(screen *ebiten.Image, number float32) {
	var (
		x float32 = constant.BOARD_BASE_POSITION_X + 1
		y float32 = constant.BOARD_BASE_POSITION_Y + (constant.BLOCK_SIZE * number) +
			(number+1)*1
		set float32 = constant.BLOCK_SIZE + 1
	)
	// Column 1
	for i := 0; i < 9; i++ {
		vector.DrawFilledRect(screen, x+(set*float32(i)), y, constant.BLOCK_SIZE,
			constant.BLOCK_SIZE, _color.BoardColor, true)
	}
}
