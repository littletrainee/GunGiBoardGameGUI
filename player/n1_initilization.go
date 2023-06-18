package player

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/pair"
)

func Initilization(selfColor color.Gray16, isown bool) Player {
	return Player{
		SelfColor: selfColor,
		IsOwn:     isown,
		KomaDai:   []pair.Pair[koma.Koma, int]{},
	}
}
