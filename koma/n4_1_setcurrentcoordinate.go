package koma

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// SetCurrentCoordinate 設定駒繪製在畫面上的位置
//
// 參數coordinate為駒要繪製的位置，shift為在目標區塊上是否有堆疊的偏移量
func (k *Koma) SetCurrentCoordinate(coordinate image.Point, shift int) {
	k.CurrentCoordinate.X = int(coordinate.X+constant.BLOCK_SIZE/2) - shift
	k.CurrentCoordinate.Y = int(coordinate.Y+constant.BLOCK_SIZE/2) - shift
}
