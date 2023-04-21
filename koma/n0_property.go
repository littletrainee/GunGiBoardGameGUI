package koma

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
	"github.com/littletrainee/pair"
)

type Koma struct {
	Name              string
	Color             color.Gray16
	CurrentCoordinate coordinate.Coordinate
	CurrentPosition   pair.Pair[int, int]
	Img               *ebiten.Image
	Op                *ebiten.DrawImageOptions
	ProbablyMoveing   [][]coordinate.Coordinate
}
