package BoardAndKomaDai

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/Color"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
)

func Board(screen *ebiten.Image) {
	// Base
	vector.DrawFilledRect(screen, Const.BOARD_BASE_POSITION_X,
		Const.BOARD_BASE_POSITION_Y, Const.BOARD_BASE_SIZE_WIDTH, Const.BOARD_BASE_SIZE_WIDTH,
		color.Black, true)
	// Row 1
	for i := 0; i < 9; i++ {
		row(screen, float32(i))
	}
	// row2(screen)
}

func row(screen *ebiten.Image, number float32) {
	var (
		x float32 = Const.BOARD_BASE_POSITION_X + 1
		y float32 = Const.BOARD_BASE_POSITION_Y + (Const.BLOCK_SIZE * number) +
			(number+1)*1
		set float32 = Const.BLOCK_SIZE + 1
	)
	// Column 1
	for i := 0; i < 9; i++ {
		vector.DrawFilledRect(screen, x+(set*float32(i)), y, Const.BLOCK_SIZE,
			Const.BLOCK_SIZE, Color.BoardColor, true)
	}
}
