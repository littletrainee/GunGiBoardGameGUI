package player

import (
	"image/color"
)

// Initilization 初始化玩家物件，設定好並回傳
func Initilization(selfColor color.Gray16, isown bool) Player {
	return Player{
		SelfColor: selfColor,
		IsOwn:     isown,
		// KomaDai:   []pair.Pair[koma.Koma, int]{},
	}
}
