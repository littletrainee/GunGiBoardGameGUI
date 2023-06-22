package level

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// OnLevelBlock 檢查滑鼠是否有按再特定等級按鈕上
//
// 若有在按鈕上則回傳true，否則false
//
// 參數x與y為滑鼠的座標
func (l *Level) OnLevelBlock(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(l.Coor.X, l.Coor.Y,
		l.Coor.X+constant.LEVEL_BLOCK_SIZE-2,
		l.Coor.Y+constant.LEVEL_BLOCK_SIZE-2))
}
