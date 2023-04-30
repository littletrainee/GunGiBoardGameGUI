package koma

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
	"github.com/littletrainee/pair"
)

type Koma struct {
	Name              string
	Color             color.Gray16
	CurrentCoordinate image.Point
	CurrentPosition   pair.Pair[int, int]
	Img               *ebiten.Image
	Op                *ebiten.DrawImageOptions
	ProbablyMoveing   [][]coordinate.Coordinate
}
