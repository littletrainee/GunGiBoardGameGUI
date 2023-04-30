package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

func (g *GameState) SetRecommendedArrangement() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, v := range g.ArrangementList {
			if v.OnRecommendedArrangement(x, y) {
				// 設定
				g.RecommendedArramgement = v.Code
				g.Phase = phase.SET_TO_PREPARE_FOR_GAMING
				g.ArrangementList = nil
				break
			}
		}
	}
}
