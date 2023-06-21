package koma

import (
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (k *Koma) OnKoma(x, y float64) bool {
	var (
		tempx float64 = x - float64(k.CurrentCoordinate.X)
		tempy float64 = y - float64(k.CurrentCoordinate.Y)
	)
	d := math.Sqrt(math.Pow(tempx, 2) + math.Pow(tempy, 2))
	return d < float64(constant.KOMA_SIZE)
}
