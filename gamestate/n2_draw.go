package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"golang.org/x/image/font"
)

func (g *GameState) Draw(screen *ebiten.Image, font font.Face) {
	switch g.Phase {
	case phase.SELECT_COLOR:
		g.Color.Draw(screen, font)
	case phase.SELECT_LEVEL:
		for _, v := range g.LevelList {
			v.Draw(screen, font)
		}
	case phase.RECOMMENDED_ARRANGEMENT:
		for _, v := range g.ArrangementList {
			v.Draw(screen, font)
		}
	}
}
