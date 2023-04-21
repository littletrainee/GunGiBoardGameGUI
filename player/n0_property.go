package player

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/pair"
)

type Player struct {
	Name       string
	KomaDai    []pair.Pair[koma.Koma, int]
	coordinate pair.Pair[float32, float32]
	Color      color.Gray16
}
