package koma

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製駒在畫面上
//
// 參數screen為要繪製的畫面
func (k *Koma) Draw(screen *ebiten.Image) {
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
	screen.DrawImage(k.Img, k.Op)
}
