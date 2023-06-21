package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// Draw 繪製Board的上的區塊
//
// 參數screen為要繪製的目標視窗，參數font為區塊所擁有的駒上的文字字型
func (b *Board) Draw(screen *ebiten.Image, font font.Face) {
	for _, v := range b.Blocks {
		v.Draw(screen, font)
	}
}
