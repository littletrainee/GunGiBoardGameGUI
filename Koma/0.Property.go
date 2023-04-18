package Koma

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/Coordinate"
)

type Koma struct {
	Name              string
	Color             color.Gray16
	CurrentCoordinate Coordinate.Coordinate
	ProbablyMoveing   [][]Coordinate.Coordinate
}
