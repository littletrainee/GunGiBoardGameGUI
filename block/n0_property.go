package block

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/pair"
)

type Block struct {
	Name         pair.Pair[int, int]
	CurrentColor color.RGBA
	Coordinate   coordinate.Coordinate
	KomaStack    []koma.Koma
	BeSelect     bool
}
