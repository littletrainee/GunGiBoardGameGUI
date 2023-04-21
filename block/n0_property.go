package block

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

type Block struct {
	Current    color.RGBA
	Coordinate coordinate.Coordinate
	KomaStack  []koma.Koma
}
