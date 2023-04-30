package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// Draw each block and Koma Dai
func (b *Board) Draw(screen *ebiten.Image, font font.Face) {
	for _, v := range b.Blocks {
		v.Draw(screen, font)
	}
	KomaTai(screen)
}
