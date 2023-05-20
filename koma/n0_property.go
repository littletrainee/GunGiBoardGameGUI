package koma

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Koma struct {
	Name              string
	Color             color.Gray16
	CurrentCoordinate image.Point
	CurrentPosition   image.Point
	Img               *ebiten.Image
	Op                *ebiten.DrawImageOptions
	ProbablyMoveing   [][][]image.Point
}
