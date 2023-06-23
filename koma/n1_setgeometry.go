package koma

import (
	"math"
)

// SetGeoMetry 設定駒是否要翻轉與畫面的繪製位置
//
// 參數rotate為是否要翻轉180°
func (k *Koma) SetGeoMetry(rotate float64) {
	if rotate == math.Pi {
		k.Op.GeoM.Rotate(math.Pi)
		k.Op.GeoM.Translate(float64(k.CurrentCoordinate.X+10), float64(k.CurrentCoordinate.Y+12))
	} else {
		k.Op.GeoM.Translate(float64(k.CurrentCoordinate.X-10), float64(k.CurrentCoordinate.Y-12))
	}
}
