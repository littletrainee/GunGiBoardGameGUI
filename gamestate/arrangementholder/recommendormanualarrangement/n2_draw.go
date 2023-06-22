package recommendormanualarrangement

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製RecommendOrManualArrangement 繪製再入門或是初級階級下的推薦布陣或是手動布陣選項
//
// 參數screen為要繪製的目標視窗
func (r *RecommendOrManualArrangement) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(r.Coor.X), float32(r.Coor.Y),
		constant.ARRANGEMENT_BUTTON_SIZE,
		constant.ARRANGEMENT_BUTTON_SIZE, color.Black, true)
	vector.DrawFilledRect(screen, float32(r.Coor.X)+1, float32(r.Coor.Y)+1,
		constant.ARRANGEMENT_BUTTON_SIZE-2,
		constant.ARRANGEMENT_BUTTON_SIZE-2, r.Background, true)
	screen.DrawImage(r.Img, r.Opt)
}
