package arrangementholder

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw 繪製玩家的推薦或自訂布陣物件在畫面上
//
// 參數screen為要繪製的畫面
func (a *ArrangementHolder) Draw(screen *ebiten.Image) {
	screen.DrawImage(a.TitleImg, a.TitleOpt)
	for _, v := range a.ArrangementList {
		v.Draw(screen)
	}
}
