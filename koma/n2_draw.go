package koma

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func (k *Koma) Draw(screen *ebiten.Image, font font.Face) {
	// base for outline
	vector.DrawFilledCircle(screen,
		float32(k.CurrentCoordinate.X), float32(k.CurrentCoordinate.Y),
		constant.KOMA_SIZE, color.Black, true)
	vector.DrawFilledCircle(screen,
		float32(k.CurrentCoordinate.X), float32(k.CurrentCoordinate.Y),
		constant.KOMA_SIZE-1, k.Color, true)

	vector.DrawFilledCircle(screen,
		float32(k.CurrentCoordinate.X), float32(k.CurrentCoordinate.Y),
		constant.KOMA_SIZE-5.25, color.Black, true)
	vector.DrawFilledCircle(screen,
		float32(k.CurrentCoordinate.X), float32(k.CurrentCoordinate.Y),
		constant.KOMA_SIZE-6.25, color.White, true)
	/*
		正向文字為當前座標X-10,Y-12
		反向文字為當前座標X+10,Y+12
	*/
	screen.DrawImage(k.Img, k.Op)
}
