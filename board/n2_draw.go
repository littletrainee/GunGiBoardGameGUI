package board

import "github.com/hajimehoshi/ebiten/v2"

// Draw each block and Koma Dai
func (b *Board) Draw(screen *ebiten.Image) {
	for _, v := range b.Blocks {
		v.Draw(screen)
	}
	KomaTai(screen)
}
