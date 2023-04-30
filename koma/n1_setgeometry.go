package koma

import "math"

func (k *Koma) SetGeoMetry(rotate float64) {
	if rotate == math.Pi {
		k.Op.GeoM.Rotate(math.Pi)
		k.Op.GeoM.Translate(float64(k.CurrentCoordinate.X+10), float64(k.CurrentCoordinate.Y+12))
	} else {
		k.Op.GeoM.Translate(float64(k.CurrentCoordinate.X-10), float64(k.CurrentCoordinate.Y-12))
	}
}
