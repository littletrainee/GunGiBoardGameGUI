package timer

import (
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (t *Timer) SetGeoM(rotate bool) {
	t.Op.GeoM.Scale(2, 2)
	if rotate {
		t.Op.GeoM.Rotate(math.Pi)
		t.Op.GeoM.Translate(
			float64(t.CurrentCoordinate.X)+float64(constant.TIMER_SIZE)/6*4.5,
			float64(t.CurrentCoordinate.Y)+float64(constant.TIMER_SIZE)/6*4)
	} else {
		t.Op.GeoM.Translate(
			float64(t.CurrentCoordinate.X)+float64(constant.TIMER_SIZE)/5,
			float64(t.CurrentCoordinate.Y)+float64(constant.TIMER_SIZE)/4)
	}
}
