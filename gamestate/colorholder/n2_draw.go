package colorholder

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw 繪製玩家的顏色選擇物件在畫面上
//
// 參數screen為要繪製的畫面
func (c *ColorHolder) Draw(screen *ebiten.Image) {
	screen.DrawImage(c.TitleImg, c.TitleOpt)
	for _, v := range c.KomaList {
		v.Draw(screen)
	}
}
