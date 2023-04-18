package Player

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/Koma"
	"github.com/littletrainee/GunGiBoardGameGUI/Player/Pair"
)

type Player struct {
	Name       string
	KomaDai    []Pair.Pair[Koma.Koma, int]
	coordinate Pair.Pair[float32, float32]
	Color      color.Gray16
}
