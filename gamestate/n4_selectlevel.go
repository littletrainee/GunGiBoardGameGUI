package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// SelectLevel 檢查玩家所選則的階級並設置本局遊戲的階級
func (g *GameState) SelectLevel() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, v := range g.LevelHolder.LevelList {
			if v.OnLevelBlock(x, y) {
				g.LevelHolder.CurrentLevel = v.Code
				g.Phase = phase.SET_LEVEL
				break
			}
		}

	}
}
