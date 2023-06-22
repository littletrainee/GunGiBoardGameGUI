package board

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw 繪製Board的上的區塊
//
// 參數screen為要繪製的目標視窗
func (b *Board) Draw(screen *ebiten.Image) {
	for _, v := range b.Blocks {
		v.Draw(screen)
	}
}
