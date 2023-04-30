package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

func (g *GameState) SelectLevel() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, v := range g.LevelList {
			if v.OnLevelBlock(x, y) {
				g.Level = v.Code
				g.Phase = phase.SET_LEVEL
				break
			}
		}

	}
}
