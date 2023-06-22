package declaresumi

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// SuMiButton 檢查滑鼠位置是否有再"布陣完畢"按鈕上。
//
// 若有在按鈕上則回傳true，否則回傳false。
//
// 參數x與y為滑鼠的座標。
func (d *DeclareSuMi) SuMiButton(x, y int) bool {
	return image.Point{x, y}.In(image.Rect(constant.SUMI_BUTTON_X,
		constant.SUMI_BUTTON_Y,
		constant.SUMI_BUTTON_X+constant.SUMI_BUTTON_WIDTH,
		constant.SUMI_BUTTON_Y+constant.SUMI_BUTTON_HEIGHT))
}
