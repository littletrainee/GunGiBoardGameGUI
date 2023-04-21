package koma

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func (k *Koma) Draw(screen *ebiten.Image, font font.Face) {
	// base for outline
	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		constant.RADIUS, color.Black, true)
	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		constant.RADIUS-1, k.Color, true)

	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		constant.RADIUS-6.5, color.Black, true)
	vector.DrawFilledCircle(screen,
		k.CurrentCoordinate.X, k.CurrentCoordinate.Y,
		constant.RADIUS-7.5, color.White, true)

	text.Draw(k.Img, k.Name, font, 0, 20, color.Black)
	/*
		正向文字為當前座標X-10,Y-12
		反向文字為當前座標X+10,Y+12

	*/
	screen.DrawImage(k.Img, k.Op)
}
