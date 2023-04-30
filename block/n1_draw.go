package block

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func (b *Block) Draw(screen *ebiten.Image, font font.Face) {
	vector.DrawFilledRect(screen, b.Coordinate.X, b.Coordinate.Y, constant.BLOCK_SIZE,
		constant.BLOCK_SIZE, b.CurrentColor, true)
	for _, v := range b.KomaStack {
		v.Draw(screen, font)
	}
}
