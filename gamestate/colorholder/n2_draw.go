package colorholder

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw 繪製玩家的顏色選擇物件在畫面上
//
// 參數screen為要繪製的畫面
func (s *ColorHolder) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.TitleImg, s.TitleOpt)
	for _, v := range s.KomaList {
		v.Draw(screen)
	}
}
