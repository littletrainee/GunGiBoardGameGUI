package block

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (b *Block) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.Coordinate.X, b.Coordinate.Y, constant.BLOCK_SIZE,
		constant.BLOCK_SIZE, color.BoardColor, true)
}
