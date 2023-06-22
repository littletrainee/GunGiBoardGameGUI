// 推薦或是手動布陣物件
package recommendormanualarrangement

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
)

// 推薦或是手動布陣
type RecommendOrManualArrangement struct {
	ArrangementSelect arrangementselect.ArrangementSelect // 布陣的選擇
	Coor              image.Point                         // 布陣按鈕的座標
	Background        color.Gray16                        // 按鈕的背景
	Img               *ebiten.Image                       // 按鈕文字的圖像
	Opt               *ebiten.DrawImageOptions            // 按鈕文字圖像的參數
}
