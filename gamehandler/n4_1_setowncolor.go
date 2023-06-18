package gamehandler

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
)

// 設定自家可放置駒的範圍
func (g *Game) setConfirmPosition() {
	var (
		block block.Block
		coor  image.Point
	)
	for row := 7; row < 10; row++ {
		for column := 1; column < 10; column++ {
			coor = image.Point{X: column, Y: row}
			block = g.Board.Blocks[coor]
			if block.HasSuI() || len(block.KomaStack) == g.GameState.MaxLayer {
				block.CurrentColor = color.DenyColor
			} else {
				block.CurrentColor = color.ConfirmColor
			}
			g.Board.Blocks[coor] = block
		}
	}
}
