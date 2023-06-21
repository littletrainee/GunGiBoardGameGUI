package block

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

// Draw 繪製Block的背景與所擁有的駒
//
// 參數screen為要繪製的目標視窗
func (b *Block) Draw(screen *ebiten.Image, font font.Face) {
	vector.DrawFilledRect(screen, b.Coordinate.X, b.Coordinate.Y, constant.BLOCK_SIZE,
		constant.BLOCK_SIZE, b.CurrentColor, true)
	for _, v := range b.KomaStack {
		v.Draw(screen, font)
	}
}
