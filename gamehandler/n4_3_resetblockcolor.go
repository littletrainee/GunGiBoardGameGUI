package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
)

// resetBlockColor 重製區塊的顏色，當玩家選擇完並放置完駒後將棋盤上可以放置駒的區域變回棋盤的顏色
//
// 參數b為棋盤物件
func resetBlockColor(b *board.Board) {
	for k, tempblock := range b.Blocks {
		if tempblock.CurrentColor != color.BoardColor {
			tempblock.CurrentColor = color.BoardColor
			b.Blocks[k] = tempblock
		}
	}

}
