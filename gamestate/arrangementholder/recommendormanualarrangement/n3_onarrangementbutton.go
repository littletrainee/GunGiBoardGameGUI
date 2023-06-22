package recommendormanualarrangement

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// OnArrangementButton 檢查滑鼠是否有按在特定的布陣按鈕上。
//
// 若有在按鈕上則回傳true，否則false。
//
// 參數x與y為滑鼠的座標。
func (r *RecommendOrManualArrangement) OnArrangementButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(r.Coor.X, r.Coor.Y,
		r.Coor.X+constant.ARRANGEMENT_BUTTON_SIZE-2,
		r.Coor.Y+constant.ARRANGEMENT_BUTTON_SIZE-2))
}
