package block

import (
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

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
