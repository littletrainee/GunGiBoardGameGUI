package Block

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/Color"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
)

func (b *Block) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.Coordinate.X, b.Coordinate.Y, Const.BLOCK_SIZE,
		Const.BLOCK_SIZE, Color.BoardColor, true)
}
