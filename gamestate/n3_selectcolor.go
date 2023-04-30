package gamestate

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/level"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/selectcolor"
)

func (g *GameState) SelectColor() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, v := range g.Color.KomaList {
			if v.OnKoma(float64(x), float64(y)) {
				if v.Color == color.Black {
					g.First = color.Black
				} else {
					g.First = color.White
				}
				g.Phase = phase.SET_COLOR
				g.LevelList = level.Initilization()
				g.Color = selectcolor.SelectColor{}
				break
			}
		}
	}
}
