package block

import (
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// Shift 駒繪製在區塊上的偏移量，會因為當前區塊已有的駒的數量而返回常數偏移。
//
// 若目標區塊沒有駒則不需要偏移，有一個則偏移常數SECOND_DAN，兩個則偏移常數THIRD_DAN。
func Shift(komastack []koma.Koma) int {
	switch len(komastack) {
	case 1:
		return constant.SECOND_DAN
	case 2:
		return constant.THIRD_DAN
	default:
		return 0
	}
}
