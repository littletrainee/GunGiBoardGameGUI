// 宣告布陣完畢的物件
package declaresumi

import "github.com/hajimehoshi/ebiten/v2"

// 宣告布陣完畢
type DeclareSuMi struct {
	Show    bool                     // 顯示的切換，true為顯示，false為否
	SuMiImg *ebiten.Image            // 顯示布陣完畢文字的圖像
	SuMiOpt *ebiten.DrawImageOptions // 布陣完畢文字的繪製參數
}
