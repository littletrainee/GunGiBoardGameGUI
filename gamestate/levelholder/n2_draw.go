package levelholder

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw 繪製玩家的階級選擇物件在畫面上
//
// 參數screen為要繪製的畫面
func (l *LevelHolder) Draw(screen *ebiten.Image) {
	screen.DrawImage(l.TitleImg, l.TitleOpt)
	for _, v := range l.LevelList {
		v.Draw(screen)
	}
}
