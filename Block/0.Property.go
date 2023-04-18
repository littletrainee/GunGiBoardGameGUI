package Block

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/Coordinate"
	"github.com/littletrainee/GunGiBoardGameGUI/Koma"
)

type Block struct {
	Current    color.RGBA
	Coordinate Coordinate.Coordinate
	KomaStack  []Koma.Koma
}
