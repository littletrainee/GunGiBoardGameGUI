package koma

import (
	"image"
)

// SetCurrentPosition 設定駒在棋盤上的位置
//
// 參數p為目標位置
func (k *Koma) SetCurrentPosition(p image.Point) {
	k.CurrentPosition = p
}
