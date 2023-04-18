package Koma

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
	"golang.org/x/image/font"
)

func (k *Koma) Draw(screen *ebiten.Image, font font.Face) {
	// base for outline
	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		Const.RADIUS, color.Black, true)
	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		Const.RADIUS-1, k.Color, true)

	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		Const.RADIUS-6.5, color.Black, true)
	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		Const.RADIUS-7.5, color.White, true)
	text.Draw(screen, k.Name, font, int(k.CurrentCoordinate.X)-11, int(k.CurrentCoordinate.Y)+8, color.Black)
}
