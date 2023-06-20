package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
)

func ResetBlockColor(b *board.Board) {
	for k, tempblock := range b.Blocks {
		if tempblock.CurrentColor != color.BoardColor {
			tempblock.CurrentColor = color.BoardColor
			b.Blocks[k] = tempblock
		}
	}

}
