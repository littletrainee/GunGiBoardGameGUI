package player

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func setOwnColor(b board.Board, isact bool, GS gamestate.GameState) {
	var (
		block block.Block
		coor  image.Point
	)
	for row := 6; row < 9; row++ {
		for column := 0; column < 9; column++ {
			coor = image.Point{X: column, Y: row}
			block = b.Blocks[coor]
			if isact {
				block.CurrentColor = color.BoardColor
			} else if block.HasSuI() || len(block.KomaStack) == GS.MaxLayer {
				block.CurrentColor = color.DenyColor
			} else {
				block.CurrentColor = color.ConfirmColor
			}
			b.Blocks[coor] = block
		}
	}
}
