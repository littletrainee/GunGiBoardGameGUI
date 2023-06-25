package gamestate

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"golang.org/x/image/font"
)

// SelectColor 檢查玩家所選擇的顏色並設置本家的顏色
//
// 參數font為文字所使用的字型來源
func (g *GameState) SelectColor(font font.Face) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, v := range g.ColorHolder.KomaList {
			if v.OnKoma(float64(x), float64(y)) {
				if v.Color == color.Black {
					g.ColorHolder.Own = color.Black
				} else {
					g.ColorHolder.Own = color.White
				}
				g.DelayedChangePhaseTo(phase.SET_COLOR)
				return
			}
		}
	}
}
